package api

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/global/response"
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/models"
	"mxshop-api/user-web/proto"
	"net/http"
	"strconv"
	"time"
)

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	// 将grpc的code转换http的状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

func HandleValidatorError(c *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error": global.RemoveTopStruct(errs.Translate(global.Trans)),
	})
	return
}

// GetUserList 获取用户信息
func GetUserList(ctx *gin.Context) {

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户：%d", currentUser.ID)

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	rsp, err := global.UserSrvClient.GetUserList(context.Background(), &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList]查询【用户列表失败】")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		//data := make(map[string]interface{})

		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			Birthday: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}

		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
}

func PassWordLogin(c *gin.Context) {
	// 表单验证
	passwordLoginForm := forms.PassWordLoginForm{}

	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	if !store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	rsp, err := global.UserSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	})
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登陆失败",
				})
			}
			return
		}
	} else {
		// 只检查到用户， 没有检查密码
		if passRsp, pasErr := global.UserSrvClient.CheckPassWord(context.Background(), &proto.PasswordCheckInfo{
			PassWord:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); pasErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passRsp.Success {
				// 生成token
				j := middlewares.NewJWT()
				claims := models.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),                 // 前面的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30*5, // 150过期
						Issuer:    "imooc",
					},
				}

				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.NickName,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30*5) * 1000,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})
			}

		}
	}

}

// Register 用户注册
func Register(c *gin.Context) {
	registerForm := forms.RegisterForm{}

	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	// 验证码校验
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	value, err := rdb.Get(context.Background(), registerForm.Mobile).Result()
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "验证码错误",
		})
		return
	} else {
		if value != registerForm.Code {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "验证码错误",
			})
			return
		}
	}

	user, err := global.UserSrvClient.CreateUser(context.Background(), &proto.CreateUserInfo{
		NickName: registerForm.Mobile,
		PassWord: registerForm.PassWord,
		Mobile:   registerForm.Mobile,
	})
	if err != nil {
		zap.S().Errorf("[Register]查询【新建用户失败】:%s", err.Error())
		HandleGrpcErrorToHttp(err, c)
		return
	}

	j := middlewares.NewJWT()
	claims := models.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),                 // 前面的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30*5, // 150过期
			Issuer:    "imooc",
		},
	}

	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.Id,
		"nick_name":  user.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30*5) * 1000,
	})

}

// PassWordLogin 登录
//func PassWordLogin(c *gin.Context) {
//	// 表单验证
//	passwordLoginForm := forms.PassWordLoginForm{}
//
//	if err := c.ShouldBind(&passwordLoginForm); err != nil {
//		HandleValidatorError(c, err)
//		return
//	}
//	userConn, err := grpc.Dial(":50051", grpc.WithInsecure())
//	if err != nil {
//		zap.S().Errorw("【GetUserList】链接用户服务失败：", "msg", err.Error())
//		c.JSON(http.StatusInternalServerError, gin.H{"msg": "用户服务连接失败"})
//		return
//	}
//	defer userConn.Close()
//	userClient := proto.NewUserClient(userConn)
//	rsp, err := userClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
//		Mobile: passwordLoginForm.Mobile,
//	})
//	fmt.Printf("获取的手机号:", passwordLoginForm.Mobile)
//	if err != nil {
//		zap.S().Errorw("【GetUserByMobile】gRPC 请求失败：", "msg", err.Error())
//
//		if e, ok := status.FromError(err); ok {
//			switch e.Code() {
//			case codes.NotFound:
//				c.JSON(http.StatusBadRequest, map[string]string{
//					"mobile": "用户不存在",
//				})
//			default:
//				c.JSON(http.StatusInternalServerError, map[string]string{
//					"mobile": "登陆失败",
//				})
//			}
//			return
//		}
//	}
//	// 生成token
//	j := middlewares.NewJWT()
//	claims := models.CustomClaims{
//		ID:          uint(rsp.Id),
//		NickName:    rsp.NickName,
//		AuthorityId: uint(rsp.Role),
//		StandardClaims: jwt.StandardClaims{
//			NotBefore: time.Now().Unix(),               // 前面的生效时间
//			ExpiresAt: time.Now().Unix() + 60*60*24*30, // 30天过期
//			Issuer:    "imooc",
//		},
//	}
//
//	token, err := j.CreateToken(claims)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{
//			"msg": "生成token失败",
//		})
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"id":         rsp.Id,
//		"nick_name":  rsp.NickName,
//		"token":      token,
//		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
//	})
//
//}
