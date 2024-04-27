package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nandonyata/10xers-test/config"
	"github.com/nandonyata/10xers-test/usecase"
)

var PORT = "3000"

func main() {
	config.ConnectDB()

	app := fiber.New()
	userService := usecase.UserService{Database: config.DB}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/user/register", userService.Register)

	fmt.Printf("Listening on port: " + PORT)

	if err := app.Listen(fmt.Sprintf(":+%s", PORT)); err != nil {
		log.Panicf("Error listening: %+v", err)
	}
}
