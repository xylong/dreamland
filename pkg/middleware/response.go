package middleware

import (
	"dreamland/pkg"
	"dreamland/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"runtime"
)

func ResponseHandler(c *gin.Context) {
	c.Next()

	code, ok := c.Get("code")
	if !ok {
		code = pkg.SUCCESS
	}

	data, ok := c.Get("data")
	if !ok {
		data = gin.H{}
	}

	c.JSON(c.Writer.Status(), gin.H{
		"code": code,
		"msg":  pkg.GetMsg(code.(int)),
		"data": data,
	})

}

func Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			buf := make([]byte, 65536)
			buf = buf[:runtime.Stack(buf, false)]
			e, ok := err.(*pkg.Error)
			if ok {
				// 翻译
				errs, ok := e.Err.(validator.ValidationErrors)
				if ok {
					trans, ok := c.Value("trans").(ut.Translator)
					if !ok {
						trans, _ = util.Uni.GetTranslator("zh")
					}
					for _, item := range errs {
						e.Msg = item.Translate(trans)
						break
					}
				}

				if e.Code >= http.StatusInternalServerError {
					log.Printf("%s\n%s", err, buf)
				}

				c.AbortWithStatusJSON(e.Code, gin.H{
					"code": e.Code,
					"msg":  e.Msg,
					"data": gin.H{},
				})

				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"msg": "服务出错，请稍后尝试",
			})
		}
	}()
	c.Next()
}
