package service

import (
	"context"
	"time"

	"github.com/BoomTHDev/golang_clean_arch/entities"
	"github.com/BoomTHDev/golang_clean_arch/pkg/custom"
	_userModel "github.com/BoomTHDev/golang_clean_arch/pkg/user/model"
	_userRepository "github.com/BoomTHDev/golang_clean_arch/pkg/user/repository"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type userServiceImpl struct {
	userRepository        _userRepository.UserRepository
	sessionRepository     _userRepository.SessionRepository
}

func NewUserServiceImpl(userRepository _userRepository.UserRepository, sessionRepository _userRepository.SessionRepository) UserService {
	return &userServiceImpl{userRepository: userRepository, sessionRepository: sessionRepository}
}

func (s *userServiceImpl) Register(registerInput *_userModel.RegisterInput) (*_userModel.RegisterResult, *custom.AppError) {
	if registerInput.Name == "" || registerInput.Email == "" || registerInput.Password == "" {
		return nil, custom.ErrInvalidInput("Invalid input provided.", nil)
	}

	_, err := s.userRepository.GetUserByEmail(registerInput.Email)
	if err != nil {
		if custom.IsDuplicateKeyError(err) {
			return nil, custom.ErrInvalidInput("Email already exists.", nil)
		}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerInput.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, custom.ErrIntervalServer("Failed to generate hashed password.", nil)
	}

	newUser := &entities.User{
		ID:       uuid.New(),
		Name:     registerInput.Name,
		Email:    registerInput.Email,
		Password: string(hashedPassword),
		Role:     entities.USER,
	}

	if err := s.userRepository.CreateUser(newUser); err != nil {
		if custom.IsDuplicateKeyError(err) {
			return nil, custom.ErrInvalidInput("Email is already exists.", nil)
		}
		return nil, custom.ErrIntervalServer("Failed to create user.", err)
	}

	sessionID := uuid.New().String()
	sessionExpiration := time.Hour * 24 * 7
	if err := s.sessionRepository.SetSession(context.Background(), sessionID, _userRepository.UserSession{
		ID:   newUser.ID.String(),
		Role: string(newUser.Role),
	}, sessionExpiration); err != nil {
		return nil, custom.ErrIntervalServer("Failed to create session.", err)
	}

	return _userModel.ToRegisterResult(newUser), nil
}
