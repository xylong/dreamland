package dto

import (
	"dreamland/pkg/util"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
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

func (r *RegisterRequest) Verify(c *gin.Context) error {
	if err := c.ShouldBind(r); err != nil {
		return err
	}
	v := c.Value("trans")
	trans, ok := v.(ut.Translator)
	if !ok {
		trans, _ = util.Uni.GetTranslator("zh")
	}
	err := util.Validate.Struct(r)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		sliceErrs := []string{}
		for _, e := range errs {
			sliceErrs = append(sliceErrs, e.Translate(trans))
		}
		return errors.New(strings.Join(sliceErrs, ","))
	}
	return nil
}
