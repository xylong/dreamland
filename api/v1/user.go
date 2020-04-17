package v1

import (
	"dreamland/pkg"
	"dreamland/pkg/db"
	"dreamland/pkg/dto"
	"dreamland/pkg/service"
	"dreamland/pkg/util"
	"fmt"
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
	key := fmt.Sprintf("user:%s", claims.ID)

	if !db.HExists(key, "id") {
		db.BatchHashSet(key, map[string]interface{}{
			"id":   claims.ID,
			"name": claims.Name,
			"exp":  claims.ExpiresAt,
		})
	}
	profile, _ := db.BatchHashGet(key, "id", "name", "esp")

	c.Set("data", profile)
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
