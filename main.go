package main

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
	"shoe-store/controller"
	"shoe-store/service"
)

func main() {
	app := fiber.New()

	db, err := sql.Open("postgres", "postgres://shoeuser:shoepass@localhost:5432/shoestore?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	shoeService := service.NewShoeService(db)
	shoeController := controller.NewShoeController(shoeService)
	shoeController.RegisterRoutes(app)

	listenErr := app.Listen(":23236")
	if listenErr != nil {
		log.Fatal(listenErr)
	}
}
