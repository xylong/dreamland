package util

import (
	en2 "github.com/go-playground/locales/en"
	zh2 "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
)

func InitValidate() {
	en := en2.New()
	zh := zh2.New()
	Uni = ut.New(en, zh)
	Validate = validator.New()
}
