package middleware

import (
	"dreamland/pkg"
	"dreamland/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "未认证",
			})
			c.Abort()
			return
		}
		claims, err := util.NewJWT().Parse(token)
		if err != nil {
			pkg.PanicError(http.StatusUnauthorized, err)
			c.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "授权过期",
			})
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
