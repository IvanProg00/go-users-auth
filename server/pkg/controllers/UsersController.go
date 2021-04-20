package controllers

import (
	"fmt"
	"net/http"
	"users-authentication/pkg/api_struct"
	"users-authentication/pkg/configs"
	"users-authentication/pkg/database"
	"users-authentication/pkg/error_utils"
	"users-authentication/pkg/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(ctx *fiber.Ctx) error {
	cursor, err := database.MI.DB.
		Collection(database.CollectionUsers).
		Find(ctx.Context(), bson.D{})
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return api_struct.ErrorMessage(ctx, error_utils.InternalServerError)
	}

	var users []models.UserShowModel
	if err := cursor.All(ctx.Context(), &users); err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return api_struct.ErrorMessage(ctx, error_utils.InternalServerError)
	}

	ctx.SendStatus(http.StatusOK)
	return api_struct.SuccessMessage(ctx, users)
}

func GetUser(ctx *fiber.Ctx) error {
	user := ctx.Locals(configs.LocalUser).(models.UserModel)
	fmt.Println(user.Email)

	ctx.SendStatus(http.StatusOK)
	return api_struct.SuccessMessage(ctx, models.FromUserToShow(user))
}

func CreateUser(ctx *fiber.Ctx) error {
	user := ctx.Locals(configs.LocalUser).(models.UserModel)

	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), configs.PasswordCost)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(error_utils.InternalServerError)
	}

	user.Password = string(password)
	res, err := database.MI.DB.
		Collection(database.CollectionUsers).
		InsertOne(ctx.Context(), user)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return api_struct.ErrorMessage(ctx, error_utils.NotAddedToDatabase)
	}

	id, _ := res.InsertedID.(primitive.ObjectID)
	user.ID = id

	userShow := models.FromUserToShow(user)

	ctx.SendStatus(http.StatusCreated)
	return api_struct.SuccessMessage(ctx, userShow)
}

func UpdateUser(ctx *fiber.Ctx) error {
	var user models.UserUpdateModel
	ctx.BodyParser(&user)
	userDB := ctx.Locals(configs.LocalUser).(models.UserModel)

	user.ID = userDB.ID

	collection := database.MI.DB.
		Collection(database.CollectionUsers)

	err := bcrypt.
		CompareHashAndPassword(
			[]byte(userDB.Password),
			[]byte(user.Password),
		)
	if user.Password != "" && err != nil {
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), configs.PasswordCost)
		if err != nil {
			ctx.SendStatus(http.StatusInternalServerError)
			return api_struct.ErrorMessage(ctx, error_utils.InternalServerError)
		}
		user.Password = string(password)
	} else {
		user.Password = userDB.Password
	}

	err = collection.
		FindOneAndUpdate(ctx.Context(), bson.M{"_id": user.ID}, bson.M{"$set": user}).
		Decode(&user)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return api_struct.ErrorMessage(ctx, error_utils.InternalServerError)
	}

	var userShow models.UserShowModel
	err = collection.
		FindOne(ctx.Context(), bson.M{"_id": user.ID}).
		Decode(&userShow)
	if err != nil {
		ctx.SendStatus(http.StatusNotFound)
		return api_struct.ErrorMessage(ctx, error_utils.UserNotFound)
	}

	ctx.SendStatus(http.StatusOK)
	return api_struct.SuccessMessage(ctx, userShow)
}

func DeleteUser(ctx *fiber.Ctx) error {
	user := ctx.Locals(configs.LocalUser).(models.UserModel)

	var userShow models.UserShowModel
	err := database.MI.DB.
		Collection(database.CollectionUsers).
		FindOneAndDelete(ctx.Context(), bson.M{"_id": user.ID}).Decode(&userShow)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return api_struct.ErrorMessage(ctx, error_utils.InternalServerError)
	}
	fmt.Println("OK")

	ctx.SendStatus(http.StatusOK)
	return api_struct.SuccessMessage(ctx, userShow)
}
