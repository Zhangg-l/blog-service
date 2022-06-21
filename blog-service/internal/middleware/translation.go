package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"

	ut "github.com/go-playground/universal-translator"
	validat "github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
)

// 翻译器
func Translation() gin.HandlerFunc {
	return func(c *gin.Context) {
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())
		locale := c.GetHeader("locale")

		trans, _ := uni.GetTranslator(locale)
		v, ok := binding.Validator.Engine().(*validat.Validate)
		if ok {
			switch locale {
			case "zh":
				_ = zh_trans.RegisterDefaultTranslations(v, trans)
			case "en":
				_ = en_trans.RegisterDefaultTranslations(v, trans)
			default:
				_ = zh_trans.RegisterDefaultTranslations(v, trans)
			}
			c.Set("trans", trans)
		}
		c.Next()
	}
}
