package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nandonyata/10xers-test/model"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader != "Bearer token123" {
		return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
			Code:    http.StatusNotFound,
			Message: "Unauthorized",
			Data:    nil,
			Error:   "",
		})
	}

	// c.Locals("user", 12) //set header

	return c.Next()
}
