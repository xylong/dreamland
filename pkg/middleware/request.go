package middleware

import (
	"dreamland/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
)

func Verify() gin.HandlerFunc {
	return func(c *gin.Context) {
		v := c.Value("trans")
		trans, ok := v.(ut.Translator)
		if !ok {
			trans, _ = util.Uni.GetTranslator("zh")
		}
	}
}
