package controllers

import (
	"encoding/json"
	"net/http"
	"users-authentication/pkg/api_struct"
	"users-authentication/pkg/configs"
	"users-authentication/pkg/models"
	"users-authentication/pkg/util"

	"github.com/cristalhq/jwt/v3"
	"github.com/gofiber/fiber/v2"
)

func ParseJWTController(ctx *fiber.Ctx) error {
	verifier, err := jwt.NewVerifierHS(jwt.HS256, []byte(configs.JWTSecretKey))
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return api_struct.ErrorMessage(ctx, util.InternalServerError)
	}

	var tokenBody models.TokenModel
	err = ctx.BodyParser(&tokenBody)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return api_struct.ErrorMessage(ctx, util.CantParse)
	}

	tokenStr := tokenBody.Token

	newToken, err := jwt.ParseAndVerifyString(tokenStr, verifier)
	if err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return api_struct.ErrorMessage(ctx, util.IncorrectJWT)
	}

	var newClaims jwt.StandardClaims
	err = json.Unmarshal(newToken.RawClaims(), &newClaims)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return api_struct.ErrorMessage(ctx, util.InternalServerError)
	}

	id := newClaims.ID

	ctx.Locals("id", id)
	return ctx.Next()
}
