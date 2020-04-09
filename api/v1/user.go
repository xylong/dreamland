package v1

import (
	"dreamland/pkg/util"
	"github.com/gin-gonic/gin"
)

var (
	User = UserController{}
)

type UserController struct {
}

func (u *UserController) Me(c *gin.Context) {
	claims := c.MustGet("claims").(*util.Claims)
	c.Set("user", claims.ID)
}
