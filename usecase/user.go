package usecase

import (
	"context"
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nandonyata/10xers-test/entity"
	"github.com/nandonyata/10xers-test/helpers"
	"github.com/nandonyata/10xers-test/model"
	"github.com/nandonyata/10xers-test/repository"
)

type UserService struct {
	Database *sql.DB
}

func (s *UserService) Register(c *fiber.Ctx) error {
	var dataFromBody entity.User
	c.BodyParser(&dataFromBody)

	if dataFromBody.Email == "" || dataFromBody.Name == "" || dataFromBody.Password == "" {
		return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
			Code:    http.StatusNotFound,
			Message: "",
			Data:    nil,
			Error:   "Fill All Field",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dataFromBody.Password = helpers.HashPassword([]byte(dataFromBody.Password))

	repo := repository.UserRepository{Database: s.Database}
	userId, err := repo.Register(ctx, dataFromBody)

	if err != nil {
		return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
			Code:    http.StatusNotFound,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
		Code:    http.StatusNotFound,
		Message: "User id: " + strconv.Itoa(userId),
		Data:    nil,
		Error:   "",
	})

}
