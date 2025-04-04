package initialize

import (
	"fmt"
	"mxshop_srvs/user_srv/global"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func initialize(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok { //更换gin的验证引擎
		china_trans := zh.New() //新建中文翻译
		english_trans := en.New()
		uni := ut.New(english_trans, china_trans, english_trans) //存储在容器中
		global.Translator, ok = uni.GetTranslator(locale)        //获取一个翻译器
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s)", locale)
		}
		switch locale {
		case "zh":
			zh_translations.RegisterDefaultTranslations(v, global.Translator) //注册中文翻译器
			break
		case "en":
			en_translations.RegisterDefaultTranslations(v, global.Translator)
			break
		default:
			zh_translations.RegisterDefaultTranslations(v, global.Translator) //注册中文翻译器
			break
		}
		return nil
	}
	return nil
}
