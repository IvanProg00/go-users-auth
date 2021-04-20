package routes

import (
	"users-authentication/pkg/controllers"
	"users-authentication/pkg/models"
	"users-authentication/pkg/validation"

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
	apiv1users.Get("/:id", validation.ValidateObjectID, models.GetUserById, controllers.GetUser)
	apiv1users.Post("", models.ValidateCreateUser, controllers.CreateUser)
	apiv1users.Put("/:id", validation.ValidateObjectID, models.ValidateUserUpdate, models.GetUserById, controllers.UpdateUser)
	apiv1users.Delete("/:id", validation.ValidateObjectID, models.GetUserById, controllers.DeleteUser)
}
