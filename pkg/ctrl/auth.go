package ctrl

import (
	"dreamland/pkg"
	"dreamland/pkg/validate"
	"github.com/gin-gonic/gin"
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
	c.Set("token", token)
}
