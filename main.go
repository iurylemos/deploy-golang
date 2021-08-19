package main

import (
	"deploy-golang/controllers"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func serveStatic(app *fiber.App) {
	app.Static("/", "./build")
}

func main() {

	//server
	app := fiber.New()
	//Handle Cors
	app.Use(cors.New())
	//Serve the build file
	serveStatic(app)
	//Setup Routes
	//setupRoutes(app)

	// controllers.SendMessage("Bom dia")

	// Server
	// port := "8080"
	port := os.Getenv("PORT")
	if port == "" {
		// panic("$PORT not set")
		port = "5000"
	}

	app.Post("/api/messages", controllers.SendMessageController)

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("It's here within")
	// 	rw.Write([]byte("Hello world"))
	// })

	// err := http.ListenAndServe(":"+port, nil)

	// if err != nil {
	// 	panic("error to listen server" + err.Error())
	// }

	app.Listen(fmt.Sprintf(":%v", port))
}
