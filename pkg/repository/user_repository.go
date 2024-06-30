package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"todo-app-go/pkg/dto"
)

type UserRepositoryInterface interface {
	Create(userData dto.CreateUserRequest) (dto.UserResponse, error)
	GetAll() ([]dto.UserResponse, error)
	GetById(userId int) (dto.UserResponse, error)
	Update(userId int, userData dto.UpdateUserRequest) error
	Delete(userId int) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(userData dto.CreateUserRequest) (dto.UserResponse, error) {
	var user dto.UserResponse
	query := fmt.Sprintf("INSERT INTO %s (username, email, password) values ($1, $2, $3) RETURNING id, username, email", usersTable)

	row := r.db.QueryRow(query, userData.Username, userData.Email, userData.Password)
	if err := row.Scan(&user.Id, &user.Username, &user.Email); err != nil {
		return dto.UserResponse{}, err
	}

	return user, nil
}

func (r *UserRepository) GetAll() ([]dto.UserResponse, error) {
	var users []dto.UserResponse

	query := fmt.Sprintf("SELECT id, username, email FROM %s tl", usersTable)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserRepository) GetById(userId int) (dto.UserResponse, error) {
	var user dto.UserResponse

	query := fmt.Sprintf(`SELECT id, username, email FROM %s WHERE id = $1`, usersTable)
	err := r.db.Get(&user, query, userId)

	return user, err
}

func (r *UserRepository) Update(userId int, userData dto.UpdateUserRequest) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if userData.Username != nil {
		setValues = append(setValues, fmt.Sprintf("username=$%d", argId))
		args = append(args, *userData.Username)
		argId++
	}

	if userData.Email != nil {
		setValues = append(setValues, fmt.Sprintf("email=$%d", argId))
		args = append(args, *userData.Email)
		argId++
	}

	// username=$1, email=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", usersTable, setQuery, argId)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Delete(userId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, usersTable)

	_, err := r.db.Exec(query, userId)

	return err
}
