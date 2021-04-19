package error_utils

import (
	"net/http"
	"users-authentication/pkg/api_struct"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	InternalServerError = "Internal server error"
	IncorrectObjectID   = "ObjectId is incorrect"
	NotAddedToDatabase  = "Not Added to Database"

	UserNotFound = "User not found"
)

func ValidateObjectID(id string, ctx *fiber.Ctx) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return objectID, ctx.JSON(api_struct.ErrorMessage(IncorrectObjectID))
	}
	return objectID, nil
}
