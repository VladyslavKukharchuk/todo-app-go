package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"todo-app-go/pkg/model"
)

type UserRepositoryInterface interface {
	Create(user model.CreateUserInput) (int, error)
	GetAll() ([]model.User, error)
	GetById(userId int) (model.User, error)
	Delete(userId int) error
	Update(userId int, input model.UpdateUserInput) error
}

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user model.CreateUserInput) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password) values ($1, $2, $3) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User

	query := fmt.Sprintf("SELECT tl.id, tl.name, tl.username FROM %s tl", usersTable)
	err := r.db.Select(&users, query)

	return users, err
}

func (r *UserRepository) GetById(userId int) (model.User, error) {
	var list model.User

	query := fmt.Sprintf(`SELECT tl.id, tl.name, tl.username FROM %s tl WHERE ul.user_id = $1`, usersTable)
	err := r.db.Get(&list, query, userId)

	return list, err
}

func (r *UserRepository) Delete(userId int) error {
	query := fmt.Sprintf(`DELETE FROM %s tl USING %s ul WHERE ul.user_id=$1`, usersTable)

	_, err := r.db.Exec(query, userId)

	return err
}

func (r *UserRepository) Update(userId int, input model.UpdateUserInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Name != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Name)
		argId++
	}

	if input.Username != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Username)
		argId++
	}

	// name=$1
	// username=$1
	// name=$1, username=$2
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE ul.list_id=$%d", usersTable, setQuery, argId, argId+1)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)
	return err
}
