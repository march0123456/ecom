package handler

import (
	"ecommerce/logs"
	"ecommerce/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderHandler {
	return orderHandler{orderService: orderService}
}

func (h orderHandler) CreateOrder(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Params("id"))
	request := service.OrderRequest{}
	err := c.BodyParser(request)
	if err != nil {
		logs.Error(err)
		return err
	}
	response, err := h.orderService.CreateOrder(userID, request)
	if err != nil {
		logs.Error(err)
		return err
	}

	return c.Status(201).JSON(response)

}

func (h orderHandler) GetOrders(c *fiber.Ctx) error {
	userID, _ := strconv.Atoi(c.Params("id"))
	response, err := h.orderService.GetOrders(userID)
	if err != nil {
		logs.Error(err)
		return err
	}
	return c.Status(200).JSON(response)
}

func (h orderHandler) CancelOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logs.Error(err)
		return err
	}
	err = h.orderService.CancelOrder(id)
	if err != nil {
		logs.Error(err)
		return err
	}
	c.Status(200)
	return nil

}
func (h orderHandler) GetOrder(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		logs.Error(err)
		return err
	}
	response, err := h.orderService.GetOrder(id)
	if err != nil {
		logs.Error(err)
		return err
	}

	return c.Status(200).JSON(response)
}
