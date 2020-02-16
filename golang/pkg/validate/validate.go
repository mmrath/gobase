package validate

import (
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/mmrath/gobase/golang/pkg/errutil"
)

// use a single instance , it caches struct info
var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    ut.Translator
)

func init() {
	enTranslator := en.New()
	uni = ut.New(enTranslator, enTranslator)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	enTran, ok := uni.GetTranslator("en")
	if !ok {
		panic("enTranslator translation not found")
	}
	trans = enTran

	validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	err := en_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(errutil.Wrap(err, "error while registering default translations"))
	}
}

func Struct(v interface{}) error {
	err := validate.Struct(v)
	return convertError(v, err)
}

func Field(field interface{}, tag string) error {
	err := validate.Var(field, tag)
	return convertError(field, err)
}

func convertError(_ interface{}, err error) error {
	if err != nil {

		//Validation syntax is invalid
		if err, ok := err.(*validator.InvalidValidationError); ok {
			return errutil.Wrap(err, "failed during validation")
		}

		if err, ok := err.(validator.ValidationErrors); ok {
			errMap := make(map[string]string)
			for _, fe := range err {
				name := fe.Field()
				msg := fe.Translate(trans)
				errMap[name] = msg
			}
			return errutil.NewFieldErrors(errMap)
		}
		return errutil.Wrap(err, "unexpected error during validation")
	}
	return nil

}
