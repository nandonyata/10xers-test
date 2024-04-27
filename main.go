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

	userService := usecase.UserService{Database: config.DB}

	app := fiber.New()
	userApi := app.Group("/user")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	userApi.Post("/register", userService.Register)
	userApi.Post("/login", userService.Login)

	fmt.Printf("Listening on port: " + PORT)
	if err := app.Listen(fmt.Sprintf(":+%s", PORT)); err != nil {
		log.Panicf("Error listening: %+v", err)
	}
}
