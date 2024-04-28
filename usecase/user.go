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
		return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
			Code:    http.StatusBadRequest,
			Message: "",
			Data:    nil,
			Error:   "Fill All Field",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dataFromBody.Password = helpers.HashPassword([]byte(dataFromBody.Password))

	repo := repository.UserRepository{Database: s.Database}
	userId, err := repo.Insert(ctx, dataFromBody)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
			Code:    http.StatusBadRequest,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
		Code:    http.StatusCreated,
		Message: "Success Register, User id: " + strconv.Itoa(userId),
		Data:    nil,
		Error:   "",
	})

}

func (s *UserService) Login(c *fiber.Ctx) error {
	var dataFromBody entity.User
	c.BodyParser(&dataFromBody)

	if dataFromBody.Email == "" || dataFromBody.Password == "" {
		return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
			Code:    http.StatusBadRequest,
			Message: "",
			Data:    nil,
			Error:   "Fill All Field",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	repo := repository.UserRepository{Database: s.Database}
	user, err := repo.FindOne(ctx, dataFromBody)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
			Code:    http.StatusBadRequest,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	checkPassword := helpers.ComparePassword([]byte(user.Password), []byte(dataFromBody.Password))
	if !checkPassword {
		return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
			Code:    http.StatusBadRequest,
			Message: "",
			Data:    nil,
			Error:   "invalid email/password",
		})
	}

	accessToken := helpers.SignToken(uint(user.Id))

	return c.Status(http.StatusOK).JSON(model.HTTPResponse{
		Code:    http.StatusOK,
		Message: "Success Login",
		Data: struct {
			AccessToken string      `json:"accessToken"`
			User        entity.User `json:"user"`
		}{
			AccessToken: accessToken,
			User:        user,
		},
		Error: "",
	})
}
