package controller

import (
	"github.com/BoomTHDev/golang_clean_arch/pkg/custom"
	_userModel "github.com/BoomTHDev/golang_clean_arch/pkg/user/model"
	_userService "github.com/BoomTHDev/golang_clean_arch/pkg/user/service"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService _userService.UserService
}

func NewUserController(userService _userService.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(c *fiber.Ctx) error {
	input := &_userModel.RegisterInput{}

	if err := c.BodyParser(input); err != nil {
		return custom.ErrInvalidInput("Invalid input provided.", err)
	}

	user, appErr := uc.userService.Register(input)
	if appErr != nil {
		return appErr
	}

	return c.JSON(user)
}
