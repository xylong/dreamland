package middleware

import (
	"dreamland/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/go-playground/validator/v10/translations/zh"
)

// Translate 翻译中间件
func Translate() gin.HandlerFunc {
	return func(c *gin.Context) {
		locale := c.DefaultQuery("locale", "zh")
		trans, _ := util.Uni.GetTranslator(locale)
		switch locale {
		case "zh":
			zh.RegisterDefaultTranslations(util.Validate, trans)
		case "en":
			en.RegisterDefaultTranslations(util.Validate, trans)
		default:
			zh.RegisterDefaultTranslations(util.Validate, trans)
		}

		c.Set("trans", trans)
		c.Next()
	}
}
