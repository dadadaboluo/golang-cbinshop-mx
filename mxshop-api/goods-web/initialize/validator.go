package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"mxshop-api/goods-web/global"
	"reflect"
	"strings"

	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok { //更换gin的验证引擎
		//将struct字段转为tag的json字段
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		zhT := zh.New() //新建中文翻译
		enT := en.New()
		uni := ut.New(enT, zhT, enT) //存储在容器中
		//获取一个翻译器
		global.Trans, ok = uni.GetTranslator(locale)

		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, global.Trans) //注册中文翻译器
			break
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.Trans)
			break
		default:
			zh_translations.RegisterDefaultTranslations(v, global.Trans) //注册中文翻译器
			break
		}
		//注册自定义验证方法

		//return nil
		return
	}
	return
	//return nil
}
