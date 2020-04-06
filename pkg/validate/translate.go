package validate

import (
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	Validate *validator.Validate
	trans    ut.Translator
)

func init() {
	ZH := zh.New()
	uni := ut.New(ZH)
	trans, _ = uni.GetTranslator("zh")
	Validate = validator.New()
	zh_translations.RegisterDefaultTranslations(Validate, trans)
}
