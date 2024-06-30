package service

import (
	"crypto/sha1"
	"fmt"
	"todo-app-go/pkg/dto"
	"todo-app-go/pkg/repository"
)

const (
	salt = "hjqrhjqw124617ajfhajs"
)

type UserServiceInterface interface {
	Create(userData dto.CreateUserRequest) (dto.UserResponse, error)
	GetAll() ([]dto.UserResponse, error)
	GetById(userId int) (dto.UserResponse, error)
	Update(userId int, userData dto.UpdateUserRequest) (dto.UserResponse, error)
	Delete(userId int) error
}

type UserService struct {
	UserRepository repository.UserRepositoryInterface
}

func NewUserService(repository repository.UserRepositoryInterface) *UserService {
	return &UserService{UserRepository: repository}
}

func (s *UserService) Create(userData dto.CreateUserRequest) (dto.UserResponse, error) {
	userData.Password = generatePasswordHash(*userData.Password)

	return s.UserRepository.Create(userData)
}

func (s *UserService) GetAll() ([]dto.UserResponse, error) {
	return s.UserRepository.GetAll()
}

func (s *UserService) GetById(userId int) (dto.UserResponse, error) {
	return s.UserRepository.GetById(userId)
}

func (s *UserService) Update(userId int, userData dto.UpdateUserRequest) (dto.UserResponse, error) {
	if err := s.UserRepository.Update(userId, userData); err != nil {
		return dto.UserResponse{}, err
	}

	return s.UserRepository.GetById(userId)
}

func (s *UserService) Delete(userId int) error {
	return s.UserRepository.Delete(userId)
}

func generatePasswordHash(password string) *string {
	hash := sha1.New()
	hash.Write([]byte(password))

	hashPassword := fmt.Sprintf("%x", hash.Sum([]byte(salt)))

	return &hashPassword
}
