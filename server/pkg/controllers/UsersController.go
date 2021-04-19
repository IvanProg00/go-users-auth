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
		return ctx.JSON(api_struct.ErrorMessage(error_utils.InternalServerError))
	}

	var users []models.UserShowModel
	if err := cursor.All(ctx.Context(), &users); err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(api_struct.ErrorMessage(error_utils.InternalServerError))
	}

	ctx.SendStatus(http.StatusOK)
	return ctx.JSON(api_struct.SuccessMessage(users))
}

func GetUser(ctx *fiber.Ctx) error {
	var user models.UserModel

	id := ctx.Params("id", "")
	user, err := models.GetUserById(id, ctx)
	if err != nil {
		return err
	}

	ctx.SendStatus(http.StatusOK)
	return ctx.
		JSON(api_struct.SuccessMessage(models.FromUserToShow(user)))
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
		return ctx.JSON(api_struct.ErrorMessage(error_utils.NotAddedToDatabase))
	}

	id, _ := res.InsertedID.(primitive.ObjectID)
	user.ID = id

	userShow := models.FromUserToShow(user)

	ctx.SendStatus(http.StatusCreated)
	return ctx.JSON(api_struct.SuccessMessage(userShow))
}

func UpdateUser(ctx *fiber.Ctx) error {
	var user models.UserUpdateModel

	ctx.BodyParser(&user)

	id := ctx.Params("id", "")
	userDB, err := models.GetUserById(id, ctx)
	if err != nil {
		return err
	}

	user.ID = userDB.ID

	collection := database.MI.DB.
		Collection(database.CollectionUsers)

	if jsonErr := models.ValidateUserUpdate(user); jsonErr != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(jsonErr)
	}

	err = bcrypt.
		CompareHashAndPassword(
			[]byte(userDB.Password),
			[]byte(user.Password),
		)
	if user.Password != "" && err != nil {
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), configs.PasswordCost)
		if err != nil {
			ctx.SendStatus(http.StatusInternalServerError)
			return ctx.JSON(api_struct.ErrorMessage(error_utils.InternalServerError))
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
		return ctx.JSON(api_struct.ErrorMessage(error_utils.InternalServerError))
	}

	var userShow models.UserShowModel
	err = collection.
		FindOne(ctx.Context(), bson.M{"_id": user.ID}).
		Decode(&userShow)
	if err != nil {
		ctx.SendStatus(http.StatusNotFound)
		return ctx.JSON(api_struct.ErrorMessage(error_utils.UserNotFound))
	}

	ctx.SendStatus(http.StatusOK)
	return ctx.JSON(api_struct.SuccessMessage(userShow))
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id", "")

	user, err := models.GetUserById(id, ctx)
	if err != nil {
		return err
	}

	var userShow models.UserShowModel
	err = database.MI.DB.
		Collection(database.CollectionUsers).
		FindOneAndDelete(ctx.Context(), bson.M{"_id": user.ID}).Decode(&userShow)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(error_utils.InternalServerError)
	}
	fmt.Println("OK")

	ctx.SendStatus(http.StatusOK)
	return ctx.JSON(api_struct.SuccessMessage(userShow))
}
