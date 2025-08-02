package repository

import "github.com/BoomTHDev/golang_clean_arch/entities"

type UserRepository interface {
	GetUserById(id string) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CreateUser(user *entities.User) error
}
