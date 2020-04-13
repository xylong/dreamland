package v1

import (
	"dreamland/pkg"
	"dreamland/pkg/service"
	"dreamland/pkg/util"
	"dreamland/pkg/validate"
	"github.com/gin-gonic/gin"
)

var (
	User        = UserController{}
	userService = service.NewUserService()
)

type UserController struct {
}

func (u *UserController) Me(c *gin.Context) {
	claims := c.MustGet("claims").(*util.Claims)
	c.Set("data", claims)
}

func (u *UserController) Store(c *gin.Context) {
	var register validate.RegisterRequest
	c.Bind(&register)
	register.Check(&register)
	if err := userService.Register(&register); err != nil {
		pkg.PanicIfErr(err)
	}
}
