package service

import (
	"ecommerce/repository"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return userService{userRepo: userRepo}
}

func (s userService) SignUp(request UserRequest) (*UserResponse, error) {
	if request.Username == "" || request.Password == "" {
		return nil, fiber.ErrExpectationFailed
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user := repository.User{
		Username: request.Username,
		Password: string(password),
	}

	newUser, err := s.userRepo.Insert(user)
	if err != nil {
		log.Fatal(err)
	}

	response := UserResponse{
		UserID:   newUser.UserID,
		Username: newUser.Username,
		Password: newUser.Password,
	}
	return &response, nil

}

func (s userService) Login(request UserRequest) (*UserResponse, error) {

	if request.Username == "" || request.Password == "" {
		return nil, fiber.ErrExpectationFailed
	}

	user, err := s.userRepo.Get(request.Username)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Incorrect username ")
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusNotFound, "Incorrect password")
	}

	response := UserResponse{
		UserID:   user.UserID,
		Username: user.Username,
		Password: user.Password,
	}

	return &response, nil

}
