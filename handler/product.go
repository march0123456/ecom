package handler

import (
	"ecommerce/logs"
	"ecommerce/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) productHandler {
	return productHandler{productService: productService}
}

func (h productHandler) NewProduct(c *fiber.Ctx) error {
	request := service.ProductRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		logs.Error(err)
		return err
	}
	response, err := h.productService.CreateProduct(request)
	if err != nil {
		logs.Error(err)
		return err
	}

	return c.Status(201).JSON(response)
}

func (h productHandler) GetAllProduct(c *fiber.Ctx) error {
	response, err := h.productService.GetAllProduct()
	if err != nil {
		logs.Error(err)
		return err
	}
	return c.Status(200).JSON(response)
}

func (h productHandler) GetProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logs.Error(err)
		return err
	}
	response, err := h.productService.GetProductById(id)
	if err != nil {
		logs.Error(err)
		return err
	}

	return c.Status(200).JSON(response)
}
