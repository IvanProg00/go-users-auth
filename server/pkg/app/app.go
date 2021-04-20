package app

import (
	"os"
	"users-authentication/pkg/database"
	"users-authentication/pkg/routes"
	"users-authentication/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run() {
	util.NewLogger()

	err := database.NewConnectionDatabase(database.DatabaseUsers)
	if err != nil {
		util.Log.Fatalln(err)
	} else {
		util.Log.Infoln(util.DatabaseConnected)
	}

	port := os.Getenv(util.Port)
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	routes.New(app)

	util.Log.Fatalln(app.Listen(":" + port))
}
