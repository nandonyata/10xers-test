package middleware

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nandonyata/10xers-test/config"
	"github.com/nandonyata/10xers-test/entity"
	"github.com/nandonyata/10xers-test/helpers"
	"github.com/nandonyata/10xers-test/model"
	"github.com/nandonyata/10xers-test/repository"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	checkToken, err := helpers.VerifyToken(authHeader)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(model.HTTPResponse{
			Code:    http.StatusUnauthorized,
			Message: "",
			Data:    nil,
			Error:   "Invalid Token",
		})
	}

	var tempUser entity.User

	idFloat, ok := checkToken["id"].(float64)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    nil,
			Error:   "Error Converting User Id",
		})
	}

	tempUser.Id = int(idFloat)

	repo := repository.UserRepository{Database: config.DB}
	user, err := repo.FindById(context.Background(), tempUser)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    nil,
			Error:   "Invalid Token",
		})
	}

	c.Locals("user", user) //set header

	return c.Next()
}
