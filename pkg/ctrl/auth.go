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
	err := validate.Validate.Struct(register)
	if err != nil {
		pkg.PanicErrorWithMsg(http.StatusBadRequest, err.Error())
	}
}
