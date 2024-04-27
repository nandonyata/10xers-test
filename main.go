package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nandonyata/10xers-test/config"
)

func main() {
	config.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
