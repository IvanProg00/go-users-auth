package controllers

import (
	"fmt"
	"users-authentication/pkg/database"
	"users-authentication/pkg/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUsers(ctx *fiber.Ctx) error {
	cursor, err := database.MI.DB.
		Collection(database.CollectionUsers).
		Find(ctx.Context(), bson.D{})
	if err != nil {
		return err
	}
	var users []models.UserModel
	cursor.All(ctx.Context(), &users)
	fmt.Println(users)

	return ctx.JSON(users)
}

func CreateUser(ctx *fiber.Ctx) error {
	var user models.UserModel
	err := ctx.BodyParser(&user)
	if err != nil {
		return err
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return err
	}

	res, err := database.MI.DB.
		Collection(database.CollectionUsers).
		InsertOne(ctx.Context(), user)
	if err != nil {
		return err
	}
	id, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return ctx.JSON(bson.M{
			"mess": "Not Found",
		})
	}
	user.ID = id
	return ctx.JSON(user)
}
