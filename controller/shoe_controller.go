package controller

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"shoe-store/service"
)

type ShoeController struct {
	ss *service.ShoeService
}

func NewShoeController(shoeService *service.ShoeService) *ShoeController {
	return &ShoeController{
		ss: shoeService,
	}
}

func (sc *ShoeController) RegisterRoutes(router fiber.Router) error {
	router.Get("/shoes", sc.GetAllShoes)

	return nil
}

func (sc *ShoeController) GetAllShoes(c *fiber.Ctx) error {
	c.Accepts(fiber.MIMEApplicationJSON)
	c.JSON(fiber.Map{
		"testing": "134",
	})
	return nil
}

func (sc *ShoeController) PostShoe(c *fiber.Ctx) error {
	log.Printf("Post Shoe %s", c.Params("shoe"))
	return nil
}
