package ctrl

import (
	"dreamland/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Authorization = AuthenticateController{}

type AuthenticateController struct {
}

func (a *AuthenticateController) Login(c *gin.Context) {
	var register validate.RegisterRequest
	c.Bind(&register)
	register.Check(register)

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": "apple",
	})
}
