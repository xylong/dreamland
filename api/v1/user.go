package v1

import (
	"dreamland/pkg"
	"dreamland/pkg/service"
	"dreamland/pkg/util"
	"dreamland/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	User        = UserController{}
	userService = service.NewUserService()
)

type UserController struct {
}

func (u *UserController) Me(c *gin.Context) {
	claims := c.MustGet("claims").(*util.Claims)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": claims,
	})
}

func (u *UserController) Store(c *gin.Context) {
	var register validate.RegisterRequest
	c.Bind(&register)
	register.Check(&register)
	token, err := userService.Register(&register)
	if err != nil {
		pkg.PanicIfErr(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "ok",
		"data": token,
	})
}
