package app

import (
	"fmt"
	"github.com/test_crud/internal/domain"
	"github.com/test_crud/internal/infra/database"
	"log"
)

type UserService interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	//FindByID(id int64) (domain.User, error)
	//Delete(id int64) error
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

func (u userService) Save(user domain.User) (domain.User, error) {
	var err error

	user.Password, err = u.passwordGen.GeneratePasswordHash(user.Password)
	if err != nil {
		return domain.User{}, fmt.Errorf("user service save user, could not generate hash: %w", err)
	}

	saveUser, err := u.userRepo.Save(user)
	if err != nil {
		return domain.User{}, fmt.Errorf("user service save user: %w", err)
	}
	return saveUser, nil
}

func (u userService) FindByEmail(email string) (domain.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		log.Println(err)
		return domain.User{}, fmt.Errorf("user service find by email user: %w", err)
	}
	return user, nil
}

//func (u userService) FindByID(id int64) (domain.User, error) {
//	user, err := u.userRepo.FindByID(id)
//	if err != nil {
//		log.Println(err)
//		return domain.User{}, fmt.Errorf("user service find by id user: %w", err)
//	}
//	return user, nil
//}

//func (u userService) Delete(id int64) error {
//	err := u.userRepo.Delete(id)
//	if err != nil {
//		return fmt.Errorf("user service delete user: %w", err)
//	}
//	return nil
//}

func (u userService) FindAll() ([]domain.User, error) {
	return []domain.User{}, nil
}
