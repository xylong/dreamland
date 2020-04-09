package ctrl

import (
	"dreamland/pkg"
	"dreamland/pkg/service"
	"dreamland/pkg/validate"
	"github.com/gin-gonic/gin"
)

var (
	userServce = service.NewUserService()
	User       = UserController{}
)

type UserController struct {
}

func (u *UserController) Register(c *gin.Context) {
	var register validate.RegisterRequest
	c.Bind(&register)
	register.Check(&register)
	token, err := userServce.Register(&register)
	if err != nil {
		pkg.PanicIfErr(err)
	}
	c.Set("token", token)
}
