package main

import (
	"deploy-golang/controllers"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

const PORT = "5000"

func serveStatic(app *fiber.App) {
	app.Static("/", "./public")
}

func main() {

	//server
	app := fiber.New()
	//Serve the build file
	serveStatic(app)

	// Server
	// port := "8080"
	// port := os.Getenv("PORT")
	port := getEnv("PORT", PORT)

	app.Post("/api/messages", controllers.SendMessageController)

	app.Listen(fmt.Sprintf(":%s", port))
}

// Gets default value passed if no value exist for given environment variable.
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
