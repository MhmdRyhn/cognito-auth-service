package validation


import (
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	// "github.com/mhmdryhn/cognito-auth-service/core/validation"
)
// CustomTranslator := ut.New(en.New(), en.New()).GetTranslator("en")


func NewValidator() *validator.Validate {
	validate := validator.New()
	RegisterCustomValidators(validate)
	// RegisterCustomErrorMessage(validate)
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
        name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
        if name == "-" {
            return ""
        }
        return name
    })
	return validate
}


// Registers custom validators
func RegisterCustomValidators(validate *validator.Validate) {
	validate.RegisterValidation("uppercase", func(field validator.FieldLevel) bool {
		return HasUppercase(field.Field().String())
	})
	validate.RegisterValidation("lowercase", func(field validator.FieldLevel) bool {
		return HasLowercase(field.Field().String())
	})
	validate.RegisterValidation("digit", func(field validator.FieldLevel) bool {
		return HasDigit(field.Field().String())
	})
	validate.RegisterValidation("punctuation", func(field validator.FieldLevel) bool {
		return HasPunctuation(field.Field().String())
	})
}


func ToCustomErrorMessage(err error) map[string]string {
	var errorMessage map[string]string = make(map[string]string)
	if fieldErrors, ok := err.(validator.ValidationErrors); ok {
        for _, err := range fieldErrors {
            switch err.Tag() {
			case "required":
				errorMessage[err.Field()] = fmt.Sprintf("%s is a required field.", err.Field())
			case "uppercase":
				errorMessage[err.Field()] = fmt.Sprintf("%s must contain at least one capital letter.", err.Field())
			case "lowercase":
				errorMessage[err.Field()] = fmt.Sprintf("%s must contain at least one small letter.", err.Field())
			case "digit":
				errorMessage[err.Field()] = fmt.Sprintf("%s must contain at least one digit.", err.Field())
			case "punctuation":
				errorMessage[err.Field()] = fmt.Sprintf("%s must contain at least one punctuation.", err.Field())
            case "password":
                errorMessage[err.Field()] = fmt.Sprintf("%s has not met all the requirements.", err.Field())
            default:
                errorMessage[err.Field()] = fmt.Sprintf("Something wrong on %s; %s", err.Field(), err.Tag())
            }
        }
    }
	return errorMessage
}


func NewTranslator() (ut.Translator, bool) {
	return ut.New(en.New(), en.New()).GetTranslator("en")
}


// Register custom error messages for different validators
func RegisterCustomErrorMessage(validate *validator.Validate) {
	// translator := en.New()
	// uni := ut.New(translator, translator)
	// trans, found := uni.GetTranslator("en")
	trans, found := NewTranslator()
	if !found {
		log.Fatal("Translator not found.")
	}
	if err := en_translations.RegisterDefaultTranslations(validate, trans); err != nil {
		log.Fatal(err)
	}
	// Register error message for `required` validator
	validate.RegisterTranslation(
		"required",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("required", "{0} is a required field.", true) 
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("required", fe.Field())
			return t
		},
	)
	// Register error message for `password` validator
	validate.RegisterTranslation(
		"password",
		trans,
		func(ut ut.Translator) error {
			return ut.Add("password", "{0} is not valid.", true) 
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("password", fe.Field())
			return t
		},
	)
}
