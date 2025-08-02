package repository

import (
	"github.com/BoomTHDev/golang_clean_arch/databases"
	"github.com/BoomTHDev/golang_clean_arch/entities"
)

type userRepositoryImpl struct {
	db databases.Database
}

func NewUserRepositoryImpl(db databases.Database) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetUserById(id string) (*entities.User, error) {
	conn := r.db.ConnectionGetting()

	user := &entities.User{}
	if err := conn.First(user, id).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepositoryImpl) GetUserByEmail(email string) (*entities.User, error) {
	conn := r.db.ConnectionGetting()

	user := &entities.User{}
	if err := conn.First(user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepositoryImpl) CreateUser(user *entities.User) error {
	conn := r.db.ConnectionGetting()

	if err := conn.Create(user).Error; err != nil {
		return err
	}

	return nil
}
