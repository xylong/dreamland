package v1

import (
	"dreamland/pkg"
	"dreamland/pkg/dto"
	"dreamland/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Authorization = AuthenticateController{}

type AuthenticateController struct {
}

func (a *AuthenticateController) Login(c *gin.Context) {
	var login dto.LoginRequest
	c.Bind(&login)
	util.Validate.Struct(&login)
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
