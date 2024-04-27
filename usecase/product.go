package usecase

import (
	"database/sql"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/nandonyata/10xers-test/model"
)

type ProductService struct {
	Database *sql.DB
}

func (s *ProductService) Create(c *fiber.Ctx) error {
	return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
		Code:    http.StatusCreated,
		Message: "CREATE PRODUCT",
		Data:    nil,
		Error:   "",
	})
}
