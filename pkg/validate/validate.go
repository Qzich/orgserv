package validate

import "github.com/go-playground/validator/v10"

var instance = validator.New(validator.WithPrivateFieldValidation())

func Struct(s interface{}) error {
	return instance.Struct(s)
}
