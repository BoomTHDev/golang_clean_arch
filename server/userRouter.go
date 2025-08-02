package server

import (
	_userController "github.com/BoomTHDev/golang_clean_arch/pkg/user/controller"
	_userRepository "github.com/BoomTHDev/golang_clean_arch/pkg/user/repository"
	_userService "github.com/BoomTHDev/golang_clean_arch/pkg/user/service"
)

func (s *fiberServer) initUserRouter() {
	userRepository := _userRepository.NewUserRepositoryImpl(s.db)
	sessionRepository := _userRepository.NewSessionRepositoryRedis(s.redis)
	userService := _userService.NewUserServiceImpl(userRepository, sessionRepository)
	userController := _userController.NewUserController(userService)

	userRouter := s.app.Group("/v1/users")
	userRouter.Post("/register", userController.Register)
}
