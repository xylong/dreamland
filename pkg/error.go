package pkg

import (
	"errors"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"net/http"
)

var (
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	ZH := zh.New()
	uni := ut.New(ZH)
	validate = validator.New()
	trans, _ = uni.GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(validate, trans)
}

type Error struct {
	uni  *ut.UniversalTranslator
	Err  error
	Msg  string
	Code int
}

func NewValidator() *validator.Validate {
	return validate
}

func (e *Error) translate() {
	// 翻译
	for _, item := range e.Err.(validator.ValidationErrors) {
		e.Msg = item.Translate(trans)
		break
	}
}

// PanicIfErr 将错误直接抛出
func PanicIfErr(err error) {
	if err != nil {
		PanicError(http.StatusInternalServerError, err)
	}
}

// PanicError 抛出错误时，携带状态码
func PanicError(code int, err error) {
	e := &Error{
		Err:  err,
		Code: code,
	}
	e.translate()
	panic(e)
}

// PanicErrorWithMsg 自定义状态码和错误信息的错误
func PanicErrorWithMsg(code int, message string) {
	panic(&Error{
		Err:  errors.New(message),
		Msg:  message,
		Code: code,
	})
}
