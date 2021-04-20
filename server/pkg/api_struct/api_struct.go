package api_struct

import (
	"github.com/gofiber/fiber/v2"
)

const (
	TokenField = "token"
)

func SuccessMessage(ctx *fiber.Ctx, data interface{}) error {
	return ctx.JSON(fiber.Map{
		"ok":   true,
		"data": data,
	})
}

func ErrorMessage(ctx *fiber.Ctx, err interface{}) error {
	return ctx.JSON(fiber.Map{
		"ok":    false,
		"error": err,
	})
}
