package validator

import (
	"errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"strings"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	validate = validator.New()
	trans, _ = uni.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(validate, trans)
}

func Valid(request interface{}) error {
	err := validate.Struct(request)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		var stringArray []string
		for _, e := range errs {
			stringArray = append(stringArray, e.Translate(trans))
		}
		return errors.New(strings.Join(stringArray, ","))
	}
	return nil
}
