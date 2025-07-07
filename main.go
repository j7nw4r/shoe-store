package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"shoe-store/controller"
	"shoe-store/service"
)

func main() {
	app := fiber.New()

	shoeService := service.NewShoeService(nil)
	shoeController := controller.NewShoeController(shoeService)
	shoeController.RegisterRoutes(app)

	listenErr := app.Listen(":23236")
	if listenErr != nil {
		log.Fatal(listenErr)
	}
}
