package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nandonyata/10xers-test/entity"
	"github.com/nandonyata/10xers-test/model"
	"github.com/nandonyata/10xers-test/repository"
)

type ProductService struct {
	Database *sql.DB
}

func (s *ProductService) Create(c *fiber.Ctx) error {
	var dataFromBody entity.Product
	c.BodyParser(&dataFromBody)

	user := c.Locals("user").(entity.User)
	dataFromBody.UserId = user.Id

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repo := repository.ProductRepository{Database: s.Database}
	productId, err := repo.Insert(ctx, dataFromBody)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
		Code:    http.StatusCreated,
		Message: "Success Create Product, product id: " + strconv.Itoa(productId),
		Data:    nil,
		Error:   "",
	})
}

func (s *ProductService) FindAll(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repo := repository.ProductRepository{Database: s.Database}
	products, err := repo.FindAll(ctx)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusCreated).JSON(model.HTTPResponse{
		Code:    http.StatusCreated,
		Message: "Success Fetching Products",
		Data:    products,
		Error:   "",
	})
}

func (s *ProductService) FindById(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    "",
			Error:   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repo := repository.ProductRepository{Database: s.Database}
	product, err := repo.FindById(ctx, productId)

	if err != nil {
		return c.Status(http.StatusNotFound).JSON(model.HTTPResponse{
			Code:    http.StatusNotFound,
			Message: "",
			Data:    nil,
			Error:   "Data Not Found",
		})
	}

	return c.Status(http.StatusOK).JSON(model.HTTPResponse{
		Code:    http.StatusOK,
		Message: "Success Fetching Products",
		Data:    product,
		Error:   "",
	})
}

func (s *ProductService) UpdateById(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    "",
			Error:   err.Error(),
		})
	}

	var dataFromBody entity.Product
	c.BodyParser(&dataFromBody)

	dataFromBody.Id = productId

	if dataFromBody.Title == "" || dataFromBody.Type == "" || dataFromBody.Price == 0 || dataFromBody.Stock == 0 {
		return c.Status(http.StatusBadRequest).JSON(model.HTTPResponse{
			Code:    http.StatusBadRequest,
			Message: "",
			Data:    nil,
			Error:   "Fill All Field",
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repo := repository.ProductRepository{Database: s.Database}

	_, err = repo.FindById(ctx, dataFromBody.Id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    nil,
			Error:   "Data Not Found",
		})
	}

	_, err = repo.UpdateById(ctx, dataFromBody)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    nil,
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(model.HTTPResponse{
		Code:    http.StatusOK,
		Message: "Success Update Product",
		Data:    "",
		Error:   "",
	})
}

func (s *ProductService) DeleteById(c *fiber.Ctx) error {
	productId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    "",
			Error:   err.Error(),
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	repo := repository.ProductRepository{Database: s.Database}
	result, err := repo.DeleteById(ctx, productId)

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(model.HTTPResponse{
			Code:    http.StatusInternalServerError,
			Message: "",
			Data:    "",
			Error:   err.Error(),
		})
	}

	return c.Status(http.StatusOK).JSON(model.HTTPResponse{
		Code:    http.StatusOK,
		Message: "Success Delete Product",
		Data:    fmt.Sprintf("Rows Affected: %d", result),
		Error:   "",
	})
}
