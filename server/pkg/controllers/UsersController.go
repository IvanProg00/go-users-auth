package controllers

import (
	"fmt"
	"net/http"
	"users-authentication/pkg/database"
	"users-authentication/pkg/models"

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

	if jsonErr := models.ValidateUser(user); jsonErr != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(jsonErr)
	}

	res, err := database.MI.DB.
		Collection(database.CollectionUsers).
		InsertOne(ctx.Context(), user)
	if err != nil {
		return err
	}

	id, _ := res.InsertedID.(primitive.ObjectID)
	user.ID = id
	ctx.SendStatus(http.StatusCreated)
	return ctx.JSON(user)
}
