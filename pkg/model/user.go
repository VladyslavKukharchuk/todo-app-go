package model

import "errors"

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreateUserInput struct {
	Name     *string `json:"name"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}

func (i CreateUserInput) Validate() error {
	if i.Name == nil && i.Username == nil && i.Password == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateUserInput struct {
	Name     *string `json:"name"`
	Username *string `json:"username"`
}

func (i UpdateUserInput) Validate() error {
	if i.Name == nil && i.Username == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
