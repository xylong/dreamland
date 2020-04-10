package ctrl

import (
	"dreamland/pkg"
	"dreamland/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Authorization = AuthenticateController{}

type AuthenticateController struct {
}

func (a *AuthenticateController) Login(c *gin.Context) {
	var login validate.LoginRequest
	c.Bind(&login)
	login.Check(&login)
	token, err := userServce.Login(&login)
	if err != nil {
		pkg.PanicIfErr(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
