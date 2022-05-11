package main

import (
	"go/compute"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Post("/getValue", getValue)

	app.Static("/", "./wwwroot", fiber.Static{
		Index:         "index.html",
		CacheDuration: time.Second * 10,
	})

	app.Listen(":8003")
}

type input struct {
	Input string `json:"input"`
}

func getValue(c *fiber.Ctx) error {
	input := input{}
	err := c.BodyParser(&input)
	if err != nil {
		return err
	}

	return c.JSON(compute.Evaluate(compute.InToPost(input.Input)))
}
