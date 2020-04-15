package dto

import (
	"dreamland/pkg/util"
	"github.com/gin-gonic/gin"
)

// LoginRequest 登录验证
type LoginRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=15"`
}

// RegisterRequest 注册验证
type RegisterRequest struct {
	LoginRequest
	Name            string `form:"name" validate:"required,min=1,max=10"`
	PasswordConfirm string `form:"password_confirm" validate:"required,eqfield=Password"`
}

func (l *LoginRequest) Verify(c *gin.Context) error {
	if err := c.ShouldBind(l); err != nil {
		return err
	}
	return util.Validate.Struct(l)
}

func (r *RegisterRequest) Verify(c *gin.Context) error {
	if err := c.ShouldBind(r); err != nil {
		return err
	}
	return util.Validate.Struct(r)
}
