package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nandonyata/10xers-test/config"
	"github.com/nandonyata/10xers-test/middleware"
	"github.com/nandonyata/10xers-test/usecase"
)

var PORT string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.ConnectDB()
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
}

func main() {
	userService := usecase.UserService{Database: config.DB}
	productService := usecase.ProductService{Database: config.DB}

	app := fiber.New()
	userApi := app.Group("/user")
	productApi := app.Group("/product")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	userApi.Post("/register", userService.Register)
	userApi.Post("/login", userService.Login)
	productApi.Post("/", middleware.AuthMiddleware, productService.Create)
	productApi.Get("/", productService.FindAll)
	productApi.Get("/:id", productService.FindById)
	productApi.Patch("/:id", middleware.AuthMiddleware, productService.UpdateById)
	productApi.Delete("/:id", middleware.AuthMiddleware, productService.DeleteById)

	fmt.Printf("Listening on port: " + PORT)
	if err := app.Listen(fmt.Sprintf(":+%s", PORT)); err != nil {
		log.Panicf("Error listening: %+v", err)
	}
}
