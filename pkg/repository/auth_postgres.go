package repository

import (
	"database/sql"
	"todo-app-go/pkg/model"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	//var id int
	//query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	//
	//row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	//if err := row.Scan(&id); err != nil {
	//	return 0, err
	//}
	//
	//return id, nil
	return 0, nil
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	//var user todo.User
	//query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	//err := r.db.Get(&user, query, username, password)
	//
	//return user, err
	return model.User{}, nil
}
