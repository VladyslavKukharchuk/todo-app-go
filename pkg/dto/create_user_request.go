package dto

import "errors"

type CreateUserRequest struct {
	Username *string `json:"username" binding:"required"`
	Email    *string `json:"email" binding:"required"`
	Password *string `json:"password" binding:"required"`
}

func (i CreateUserRequest) Validate() error {
	if i.Username == nil && i.Email == nil && i.Password == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
