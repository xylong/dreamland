package middleware

import (
	"dreamland/pkg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime"
)

func ResponseHandler(c *gin.Context) {
	c.Next()

	if c.Writer.Status() != http.StatusOK {
		c.JSON(c.Writer.Status(), gin.H{
			"code": 1,
			"msg":  "",
			"data": gin.H{},
		})
	}
}

func Recovery(c *gin.Context) {
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
					"code": e.Code,
					"msg":  e.Msg,
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
