package forms

type SendSmsForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,len=11,mobile"`
	Type   string `form:"type" json:"type" binding:"required,oneof=1 2"`

	// 1.注册发送短信验证码和动态验证码登录发送验证码

}
