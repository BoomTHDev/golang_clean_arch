package service

import (
	"github.com/BoomTHDev/golang_clean_arch/pkg/custom"
	_userModel "github.com/BoomTHDev/golang_clean_arch/pkg/user/model"
)

type UserService interface {
	Register(input *_userModel.RegisterInput) (*_userModel.RegisterResult, *custom.AppError)
}
