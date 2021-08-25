package main

import (
	"os"

	"github.com/ferjoaguilar/rest-go.git/src/config"
	"github.com/ferjoaguilar/rest-go.git/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	app := fiber.New()
	config.Connection()
	routes.SetupStudentRoutes(app)
	app.Listen(os.Getenv("PORT"))
}
