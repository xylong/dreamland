package v1

import (
	"dreamland/pkg"
	"dreamland/pkg/validate"
	"errors"
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
	token, err := userService.Login(&login)
	if err != nil {
		pkg.PanicIfErr(err)
	}
	c.Set("data", map[string]string{
		"token": token,
	})
}

func (a *AuthenticateController) Current(c *gin.Context) {
	credential := c.Request.Header.Get("Authorization")
	if credential == "" {
		pkg.PanicError(http.StatusUnauthorized, errors.New("未认证"))
	}
	token, err := userService.RefreshToken(credential)
	if err != nil {
		pkg.PanicIfErr(err)
	}
	c.Set("data", map[string]string{
		"token": token,
	})
}
