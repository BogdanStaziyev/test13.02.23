package app

import (
	"fmt"
	"github.com/test_crud/internal/domain"
	"github.com/test_crud/internal/infra/database"
)

const (
	errorPasswordGenerate = "user service save user, could not generate hash"
	errorSaveUser         = "user service save user"
	errorFindByEmail      = "user service find by email user"
	errorFindAll          = "user service find all users"
)

type UserService interface {
	Save(user domain.User) error
	FindByEmail(email string) (domain.User, error)
	FindAll() ([]domain.User, error)
}

type userService struct {
	userRepo    database.UserRepo
	passwordGen Generator
}

func NewUserService(ur database.UserRepo, gs Generator) UserService {
	return userService{
		userRepo:    ur,
		passwordGen: gs,
	}
}

func (u userService) Save(user domain.User) error {
	var err error
	user.Password, err = u.passwordGen.GeneratePasswordHash(user.Password)
	if err != nil {
		return fmt.Errorf("%s: %w", errorPasswordGenerate, err)
	}
	err = u.userRepo.Save(user)
	if err != nil {
		return fmt.Errorf("%s: %w", errorSaveUser, err)
	}
	return nil
}

func (u userService) FindByEmail(email string) (domain.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return domain.User{}, fmt.Errorf("%s: %w", errorFindByEmail, err)
	}
	return user, nil
}

func (u userService) FindAll() ([]domain.User, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return []domain.User{}, fmt.Errorf("%s: %w", errorFindAll, err)
	}
	return users, nil
}
