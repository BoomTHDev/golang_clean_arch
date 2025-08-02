package model

import (
	"github.com/BoomTHDev/golang_clean_arch/entities"
	"github.com/google/uuid"
)

type (
	RegisterInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	RegisterResult struct {
		ID   uuid.UUID         `json:"id"`
		Role entities.UserRole `json:"role"`
	}
)

func ToRegisterResult(user *entities.User) *RegisterResult {
	return &RegisterResult{
		ID:   user.ID,
		Role: user.Role,
	}
}
