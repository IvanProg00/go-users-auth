package controllers

import (
	"net/http"
	"users-authentication/pkg/api_mess"
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
		return ctx.JSON(api_mess.ErrorMessage(error_utils.InternalServerError))
	}

	var users []models.UserModel
	if err := cursor.All(ctx.Context(), &users); err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(api_mess.ErrorMessage(error_utils.InternalServerError))
	}

	ctx.SendStatus(http.StatusOK)
	return ctx.JSON(api_mess.SuccessMessage(users))
}

func GetUser(ctx *fiber.Ctx) error {
	var user models.UserModel

	id := ctx.Params("id", "")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(api_mess.ErrorMessage(error_utils.IncorrectObjectID))
	}

	err = database.MI.DB.
		Collection(database.CollectionUsers).
		FindOne(ctx.Context(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		ctx.SendStatus(http.StatusNotFound)
		return ctx.JSON(api_mess.ErrorMessage(error_utils.UserNotFound))
	}

	ctx.SendStatus(http.StatusOK)
	return ctx.JSON(api_mess.SuccessMessage(user))
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
		return ctx.JSON(api_mess.ErrorMessage(error_utils.NotAddedToDatabase))
	}

	id, _ := res.InsertedID.(primitive.ObjectID)
	user.ID = id

	ctx.SendStatus(http.StatusCreated)
	return ctx.JSON(api_mess.SuccessMessage(user))
}

func UpdateUser(ctx *fiber.Ctx) error {
	var user models.UserUpdateModel
	var userDB models.UserModel

	ctx.BodyParser(&user)

	id := ctx.Params("id", "")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(api_mess.ErrorMessage(error_utils.IncorrectObjectID))
	}
	user.ID = objectID

	collection := database.MI.DB.
		Collection(database.CollectionUsers)
	err = collection.
		FindOne(ctx.Context(), bson.M{"_id": objectID}).
		Decode(&userDB)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(api_mess.ErrorMessage(error_utils.UserNotFound))
	}

	if jsonErr := models.ValidateUserUpdate(user); jsonErr != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(jsonErr)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); user.Password != "" && err != nil {
		password, err := bcrypt.GenerateFromPassword([]byte(user.Password), configs.PasswordCost)
		if err != nil {
			ctx.SendStatus(http.StatusInternalServerError)
			return ctx.JSON(api_mess.ErrorMessage(error_utils.InternalServerError))
		}
		user.Password = string(password)
	} else {
		user.Password = userDB.Password
	}

	err = collection.
		FindOneAndUpdate(ctx.Context(), bson.M{"_id": objectID}, bson.M{"$set": user}).
		Decode(&user)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return ctx.JSON(api_mess.ErrorMessage(error_utils.NotAddedToDatabase))
	}

	err = collection.
		FindOne(ctx.Context(), bson.M{"_id": objectID}).
		Decode(&userDB)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return ctx.JSON(api_mess.ErrorMessage(error_utils.UserNotFound))
	}

	ctx.SendStatus(http.StatusOK)
	return ctx.JSON(api_mess.SuccessMessage(userDB))
}
