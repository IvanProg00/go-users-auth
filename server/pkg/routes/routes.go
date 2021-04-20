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

	ApiV1UsersRouter(apiv1)
	ApiV1AuthRouter(apiv1)
}

func ApiV1UsersRouter(routes fiber.Router) {
	users := routes.Group("/users")

	users.Get("", controllers.GetUsers)
	users.Get("/:id", validation.ParseObjectIDController, models.GetUserById, controllers.GetUser)
	users.Post("", models.ParseCreateUser, controllers.CreateUser)
	users.Put("/:id", validation.ParseObjectIDController, models.ParseUserUpdate, models.GetUserById, controllers.UpdateUser)
	users.Delete("/:id", validation.ParseObjectIDController, models.GetUserById, controllers.DeleteUser)
}

func ApiV1AuthRouter(routes fiber.Router) {
	auth := routes.Group("/auth")

	auth.Post("/login", models.ParseLogin, controllers.LoginController)
	auth.Post("", controllers.ParseJWTController, validation.ParseJWTToIDController, models.GetUserById, controllers.GetUser)
}
