package router

import (
	"dreamland/pkg"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"net/http"
	"runtime"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	ZH := zh.New()
	uni := ut.New(ZH)
	trans, _ = uni.GetTranslator("zh")
	validate = validator.New()
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

func responseHandler(c *gin.Context) {
	c.Next()

	if c.Writer.Status() == http.StatusNotFound && c.Writer.Size() <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "not found",
		})
		return
	}

	if c.Writer.Status() == http.StatusOK && c.Writer.Size() <= 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
		return
	}
}

func recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 65536)
			buf = buf[:runtime.Stack(buf, false)]
			e, ok := err.(*pkg.Error)
			if ok {
				if e.Code >= http.StatusInternalServerError {
					log.Printf("%s\n%s", err, buf)
				}

				c.AbortWithStatusJSON(e.Code, gin.H{
					"msg": e.Msg,
				})

				return
			}
			log.Printf("%s\n%s", err, buf)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "服务出错，请稍后尝试",
			})
		}
	}()
	c.Next()
}
