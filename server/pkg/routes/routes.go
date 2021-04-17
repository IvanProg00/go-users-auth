package routes

import (
	"users-authentication/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

const (
	ApiV1 = "/api/v1"

	UsersUrl = "users"
)

func New(app *fiber.App) {
	app.Get(ApiV1+"/"+UsersUrl, controllers.GetUsers)
	app.Post(ApiV1+"/"+UsersUrl, controllers.CreateUser)
}
