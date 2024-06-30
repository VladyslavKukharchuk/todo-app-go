package dto

import "errors"

type UpdateUserRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

func (i UpdateUserRequest) Validate() error {
	if i.Username == nil && i.Email == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
