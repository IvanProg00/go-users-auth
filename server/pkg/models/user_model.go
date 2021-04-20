package models

import (
	"net/http"
	"users-authentication/pkg/api_struct"
	"users-authentication/pkg/configs"
	"users-authentication/pkg/database"
	"users-authentication/pkg/util"
	"users-authentication/pkg/validation"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userFields = []validation.FieldValidateModel{
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

type UserLoginModel struct {
	Username string `json:"username" bson:"username,omitempty" validate:"required,min=2,max=28"`
	Password string `json:"password" bson:"password,omitempty" validate:"required,min=8,max=22"`
}

type TokenModel struct {
	Token string `json:"token"`
}

func ParseCreateUser(ctx *fiber.Ctx) error {
	var user UserModel
	err := ctx.BodyParser(&user)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return api_struct.ErrorMessage(ctx, util.CantParse)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return api_struct.ErrorMessage(ctx, validation.ValidateModel(err, userFields))
	}

	ctx.Locals(configs.LocalUser, user)
	return ctx.Next()
}

func ParseUserUpdate(ctx *fiber.Ctx) error {
	var user UserUpdateModel
	err := ctx.BodyParser(&user)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return api_struct.ErrorMessage(ctx, util.CantParse)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return api_struct.ErrorMessage(ctx, validation.ValidateModel(err, userFields))
	}

	return ctx.Next()
}

func GetUserById(ctx *fiber.Ctx) error {
	var user UserModel
	id := ctx.Locals(configs.LocalObjectID).(primitive.ObjectID)

	err := database.MI.DB.
		Collection(database.CollectionUsers).
		FindOne(ctx.Context(), bson.M{"_id": id}).Decode(&user)
	if err != nil {
		ctx.SendStatus(http.StatusNotFound)
		return api_struct.ErrorMessage(ctx, util.UserNotFound)
	}

	ctx.Locals(configs.LocalUser, user)
	return ctx.Next()
}

func FromUserToShow(user UserModel) UserShowModel {
	return UserShowModel{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}
}

func ParseLogin(ctx *fiber.Ctx) error {
	var user UserLoginModel
	err := ctx.BodyParser(&user)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return api_struct.ErrorMessage(ctx, util.CantParse)
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return api_struct.ErrorMessage(ctx, validation.ValidateModel(err, userFields))
	}

	ctx.Locals(configs.LocalLogin, user)
	return ctx.Next()
}
