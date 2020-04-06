package validate

// RegisterRequest 注册验证
type RegisterRequest struct {
	Name            string `form:"name" validate:"required,min=1,max=10"`
	Email           string `form:"email" validate:"required,email"`
	Password        string `form:"password" validate:"required,min=6,max=15"`
	PasswordConfirm string `form:"password_confirm" validate:"required,eqfield=Password"`
}
