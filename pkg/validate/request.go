package validate

import (
	"dreamland/pkg"
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

type Request struct {
}

func (r *Request) Check(obj interface{}) {
	err := validate.Struct(obj)
	if err != nil {
		for _, item := range err.(validator.ValidationErrors) {
			pkg.PanicErrorWithMsg(http.StatusBadRequest, item.Translate(trans))
			break
		}
	}
}
