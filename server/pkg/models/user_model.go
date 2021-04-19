package models

import (
	"net/http"
	"users-authentication/pkg/api_struct"
	"users-authentication/pkg/database"
	"users-authentication/pkg/error_utils"
	"users-authentication/pkg/validate_fields"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

type UserShowModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty"`
	Email    string             `json:"email" bson:"email,omitempty"`
}

type UserModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username string             `json:"username" bson:"username,omitempty" validate:"required,min=2,max=28"`
	Email    string             `json:"email" bson:"email,omitempty" validate:"required,email"`
	Password string             `json:"password" bson:"password,omitempty" validate:"required,min=8,max=22"`
}

type UserUpdateModel struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
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

func GetUserById(id string, ctx *fiber.Ctx) (UserModel, error) {
	var user UserModel

	objectID, err := error_utils.ValidateObjectID(id, ctx)
	if err != nil {
		return user, err
	}

	err = database.MI.DB.
		Collection(database.CollectionUsers).
		FindOne(ctx.Context(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		ctx.SendStatus(http.StatusNotFound)
		return user, ctx.JSON(api_struct.ErrorMessage(error_utils.UserNotFound))
	}

	return user, nil
}

func FromUserToShow(user UserModel) UserShowModel {
	return UserShowModel{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Password,
	}
}
