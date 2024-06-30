package service

import (
	"todo-app-go/pkg/repository"
)

type Service struct {
	UserService UserServiceInterface
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repos.UserRepository),
	}
}
