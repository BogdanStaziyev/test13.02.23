package app

import (
	"errors"
	"fmt"
	"github.com/test_crud/config"
	"github.com/test_crud/internal/domain"
	"github.com/test_crud/internal/infra/http/requests"
	"github.com/upper/db/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(user domain.User) (domain.User, error)
	Login(user requests.LoginAuth) ([]domain.User, error)
}

type authService struct {
	userService UserService
	config      config.Configuration
}

func NewAuthService(us UserService, cf config.Configuration) AuthService {
	return authService{
		userService: us,
		config:      cf,
	}
}

func (a authService) Register(user domain.User) (domain.User, error) {
	_, err := a.userService.FindByEmail(user.Email)
	if err == nil {
		return domain.User{}, fmt.Errorf("auth service error register invalid credentials user exist")
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		return domain.User{}, fmt.Errorf("auth service error register")
	}
	user, err = a.userService.Save(user)
	if err != nil {
		return domain.User{}, fmt.Errorf("auth service error register save user: %w", err)
	}
	return user, nil
}

func (a authService) Login(user requests.LoginAuth) ([]domain.User, error) {
	u, err := a.userService.FindByEmail(user.Email)
	if err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			return []domain.User{}, fmt.Errorf("auth service error login, invalid credentials user not exist: %w", err)
		}
		return []domain.User{}, fmt.Errorf("auth service error login user invalid email or password: %w", err)
	}
	valid := a.checkPasswordHash(user.Password, u.Password)
	if !valid {
		return []domain.User{}, fmt.Errorf("auth service error login user invalid email or password: %w", err)
	}
	users, err := a.userService.FindAll()
	if err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			//todo change error
			return []domain.User{}, fmt.Errorf("auth service error login, invalid credentials: %w", err)
		}
	}
	return users, nil
}

func (a authService) checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
