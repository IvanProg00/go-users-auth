package api_struct

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func SuccessMessage(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(bson.M{
		"ok":   true,
		"data": data,
	})
}

func ErrorMessage(ctx *fiber.Ctx, err interface{}) error {
	return ctx.JSON(bson.M{
		"ok":    false,
		"error": err,
	})
}
