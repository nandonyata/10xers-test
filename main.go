package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nandonyata/10xers-test/config"
)

var PORT = "3000"

func main() {
	config.ConnectDB()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(fmt.Sprintf(":+%s", PORT))
}
