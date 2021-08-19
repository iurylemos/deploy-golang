package main

import (
	"deploy-golang/controllers"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	app := fiber.New()
	app.Use(cors.New())

	// controllers.SendMessage("Bom dia")

	// Server
	// port := "8080"
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT not set")
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

	log.Fatal(app.Listen(fmt.Sprintf(":%v", port)))
}
