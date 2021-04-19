package routes

import (
	"users-authentication/pkg/controllers"

	"github.com/gofiber/fiber/v2"
)

const (
	UsersUrl = "/users"
)

func New(app *fiber.App) {
	api := app.Group("/api")
	apiv1 := api.Group("/v1")

	apiv1users := apiv1.Group("/users")
	apiv1users.Get("", controllers.GetUsers)
	apiv1users.Get("/:id", controllers.GetUser)
	apiv1users.Post("", controllers.CreateUser)
	apiv1users.Put("/:id", controllers.UpdateUser)
	apiv1users.Delete("/:id", controllers.DeleteUser)
}
