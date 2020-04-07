package ctrl

import (
	"dreamland/pkg"
	"dreamland/pkg/validate"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Authorization = AuthenticateController{}

type AuthenticateController struct {
}

func (a *AuthenticateController) Login(c *gin.Context) {
	var register validate.RegisterRequest
	c.Bind(&register)
	err := pkg.Validate.Struct(register)
	if err != nil {
		pkg.PanicError(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{
		"data": "register",
	})
}
