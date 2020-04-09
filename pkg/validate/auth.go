package validate

// LoginRequest 登录验证
type LoginRequest struct {
	Request
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=15"`
}

// RegisterRequest 注册验证
type RegisterRequest struct {
	LoginRequest
	Name            string `form:"name" validate:"required,min=1,max=10"`
	PasswordConfirm string `form:"password_confirm" validate:"required,eqfield=Password"`
}
