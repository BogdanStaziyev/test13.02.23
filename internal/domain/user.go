package domain

import (
	"github.com/test_crud/internal/infra/http/response"
	"time"
)

type User struct {
	ID          string
	Email       string
	Name        string
	Password    string
	CreatedDate time.Time
	UpdatedDate time.Time
}

func (u User) DomainToResponse() response.UserResponse {
	return response.UserResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}
