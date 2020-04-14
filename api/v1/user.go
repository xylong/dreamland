package v1

import (
	"dreamland/pkg"
	"dreamland/pkg/dto"
	"dreamland/pkg/service"
	"dreamland/pkg/util"
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
	c.Set("data", claims)
}

func (u *UserController) Store(c *gin.Context) {
	req := dto.RegisterRequest{}
	if err := req.Verify(c); err != nil {
		pkg.PanicError(http.StatusBadRequest, err)
	}
	if err := userService.Register(&req); err != nil {
		pkg.PanicIfErr(err)
	}
}
