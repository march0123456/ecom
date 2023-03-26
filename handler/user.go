package handler

import (
	"ecommerce/logs"
	"ecommerce/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) userHandler {
	return userHandler{userService: userService}
}

const jwtSecret = "secret"

func (h userHandler) SignUp(c *fiber.Ctx) error {
	request := service.UserRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		return err
	}

	if request.Username == "" || request.Password == "" {
		return fiber.ErrUnprocessableEntity
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	response, err := h.userService.SignUp(service.UserRequest{
		Username: request.Username,
		Password: string(password),
	})
	if err != nil {
		logs.Error(err)
		return err
	}

	return c.Status(201).JSON(response)
}

func (h userHandler) Login(c *fiber.Ctx) error {

	request := service.UserRequest{}
	err := c.BodyParser(&request)
	if err != nil {
		logs.Error(err)
		return fiber.ErrBadRequest
	}

	if request.Username == "" || request.Password == "" {

		return fiber.ErrUnprocessableEntity
	}
	response, err := h.userService.Login(request)
	if err != nil {
		logs.Error(err)
		return fiber.ErrInternalServerError

	}
	cliams := jwt.RegisteredClaims{
		Issuer: strconv.Itoa(response.UserID),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		logs.Error(err)
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{"jwtToken": token})
}
