package validate_fields

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

type FieldValidateModel struct {
	JSONField  string
	ModelField string
}

func ValidateFieldError(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "field is required"
	case "min":
		return fmt.Sprintf("Minimum %s symbols", err.Param())
	case "max":
		return fmt.Sprintf("Maximum %s symbols", err.Param())
	case "email":
		return "email is incorrect"
	}
	return err.Error()
}

func ValidateModel(err error, modelFields []FieldValidateModel) bson.M {
	errs := bson.M{}
	for _, err := range err.(validator.ValidationErrors) {
		for _, field := range modelFields {
			if err.Field() == field.ModelField {
				errs[field.JSONField] = ValidateFieldError(err)
			}
		}
	}
	return errs
}
