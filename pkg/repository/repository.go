package repository

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	UserRepository UserRepositoryInterface
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		UserRepository: NewUserRepository(db),
	}
}
