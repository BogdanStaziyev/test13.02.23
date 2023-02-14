package requests

import (
	"github.com/test_crud/internal/domain"
)

type LoginAuth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type RegisterAuth struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
	Name     string `json:"name" validate:"required,gte=3"`
}

func (r RegisterAuth) RegisterToUser() domain.User {
	return domain.User{
		Email:    r.Email,
		Name:     r.Name,
		Password: r.Password,
	}
}
