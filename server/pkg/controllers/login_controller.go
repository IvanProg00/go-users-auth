package controllers

import (
	"net/http"
	"users-authentication/pkg/api_struct"
	"users-authentication/pkg/configs"
	"users-authentication/pkg/database"
	"users-authentication/pkg/models"
	"users-authentication/pkg/util"

	"github.com/cristalhq/jwt/v3"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func LoginController(ctx *fiber.Ctx) error {
	user := ctx.Locals(configs.LocalLogin).(models.UserLoginModel)
	var userFound models.UserModel

	err := database.MI.DB.
		Collection(database.CollectionUsers).
		FindOne(ctx.Context(), bson.M{"username": user.Username}).
		Decode(&userFound)
	if err != nil {
		ctx.SendStatus(http.StatusNotFound)
		return api_struct.ErrorMessage(ctx, util.UserNotFound)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(user.Password)); err != nil {
		ctx.SendStatus(http.StatusBadRequest)
		return api_struct.ErrorMessage(ctx, util.UserNotFound)
	}

	signer, err := jwt.NewSignerHS(jwt.HS256, []byte(configs.JWTSecretKey))
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return api_struct.ErrorMessage(ctx, util.CreateJWTError)
	}

	claims := &jwt.RegisteredClaims{
		Audience: []string{"admin"},
		ID:       userFound.ID.Hex(),
	}
	builder := jwt.NewBuilder(signer)
	token, err := builder.Build(claims)
	if err != nil {
		ctx.SendStatus(http.StatusInternalServerError)
		return api_struct.ErrorMessage(ctx, "Error in creating token")
	}

	return api_struct.SuccessMessage(ctx, fiber.Map{api_struct.TokenField: token.String()})
}
