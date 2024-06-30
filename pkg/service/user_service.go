package service

import (
	"crypto/sha1"
	"fmt"
	"todo-app-go/pkg/model"
	"todo-app-go/pkg/repository"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

type UserServiceInterface interface {
	Create(input model.CreateUserInput) (int, error)
	GetAll() ([]model.User, error)
	GetById(userId int) (model.User, error)
	Delete(userId int) error
	Update(userId int, input model.UpdateUserInput) error
}

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) *UserService {
	return &UserService{UserRepository: repository}
}

func (s *UserService) Create(input model.CreateUserInput) (int, error) {
	input.Password = generatePasswordHash(*input.Password)

	return s.UserRepository.Create(input)
}

func (s *UserService) GetAll() ([]model.User, error) {
	return s.UserRepository.GetAll()
}

func (s *UserService) GetById(userId int) (model.User, error) {
	return s.UserRepository.GetById(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.UserRepository.Delete(userId)
}

func (s *UserService) Update(userId int, input model.UpdateUserInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.UserRepository.Update(userId, input)
}

func generatePasswordHash(password string) *string {
	hash := sha1.New()
	hash.Write([]byte(password))

	hashPassword := fmt.Sprintf("%x", hash.Sum([]byte(salt)))

	return &hashPassword
}
