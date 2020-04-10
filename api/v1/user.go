package v1

import (
	"dreamland/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	User = UserController{}
)

type UserController struct {
}

func (u *UserController) Me(c *gin.Context) {
	claims := c.MustGet("claims").(*util.Claims)
	c.JSON(http.StatusOK, gin.H{
		"profile": claims,
	})
}
