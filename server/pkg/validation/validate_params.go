package validation

import (
	"net/http"
	"users-authentication/pkg/api_struct"
	"users-authentication/pkg/configs"
	"users-authentication/pkg/error_utils"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ValidateObjectID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return api_struct.ErrorMessage(ctx, error_utils.IncorrectObjectID)
	}

	ctx.Locals(configs.LocalObjectID, objectID)
	return ctx.Next()

}
