package models

import (
	"users-authentication/pkg/validate_fields"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userFields = []validate_fields.FieldValidateModel{
	{
		JSONField:  "username",
		ModelField: "Username",
	},
	{
		JSONField:  "email",
		ModelField: "Email",
	},
	{
		JSONField:  "password",
		ModelField: "Password",
	},
}

type UserModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty" validate:""`
	Username string             `json:"username" bson:"username,omitempty" validate:"required,min=2,max=28"`
	Email    string             `json:"email" bson:"email,omitempty" validate:"required,email"`
	Password string             `json:"password" bson:"password,omitempty" validate:"required,min=8,max=22"`
}

type UserUpdateModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty" validate:""`
	Username string             `json:"username" bson:"username,omitempty" validate:"omitempty,min=2,max=28"`
	Email    string             `json:"email" bson:"email,omitempty" validate:"omitempty,email"`
	Password string             `json:"password" bson:"password,omitempty" validate:"omitempty,min=8,max=22"`
}

func ValidateUser(user UserModel) bson.M {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return validate_fields.ValidateModel(err, userFields)
	}
	return nil
}

func ValidateUserUpdate(user UserUpdateModel) bson.M {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return validate_fields.ValidateModel(err, userFields)
	}
	return nil
}
