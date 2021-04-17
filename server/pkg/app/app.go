package app

import (
	"log"
	"os"
	"users-authentication/pkg/database"
	"users-authentication/pkg/routes"
	"users-authentication/pkg/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Run() {
	err := database.NewConnectionDatabase(database.DatabaseUsers)
	if err != nil {
		log.Fatalln(err)
	}

	port := os.Getenv(util.Port)
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	app.Use(cors.New())
	app.Use(logger.New())
	routes.New(app)

	log.Fatalln(app.Listen(":" + port))
}
