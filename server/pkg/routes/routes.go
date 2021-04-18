package routes

import (
	"users-authentication/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

const (
	ApiV1 = "/api/v1"

	UsersUrl = "users"

	UsersApiV1 = ApiV1 + "/" + UsersUrl
)

func New(app *fiber.App) {

	app.Get(UsersApiV1, controllers.GetUsers)
	app.Get(UsersApiV1+"/:id", controllers.GetUser)
	app.Post(UsersApiV1, controllers.CreateUser)
	app.Put(UsersApiV1+"/:id", controllers.UpdateUser)
}
