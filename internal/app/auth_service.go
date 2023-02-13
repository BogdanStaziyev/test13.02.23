package app

import (
	"errors"
	"fmt"
	"github.com/test_crud/config"
	"github.com/test_crud/internal/domain"
	"github.com/test_crud/internal/infra/http/requests"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

const (
	ErrorRegisterUserExist = "auth service error register invalid credentials user exist"
	RegisterError          = "auth service error register"
	ErrorSave              = "auth service error register save user"
	ErrorLoginUserNotExist = "auth service error login, invalid credentials user not exist"
	ErrorLoginInvalid      = "auth service error login user invalid email or password"
)

type AuthService interface {
	Register(user domain.User) (domain.User, error)
	Login(user requests.LoginAuth) error
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
		return domain.User{}, fmt.Errorf("%s: %w", ErrorRegisterUserExist, err)
	} else if !errors.Is(err, mongo.ErrNoDocuments) {
		return domain.User{}, fmt.Errorf("%s: %w", RegisterError, err)
	}
	user, err = a.userService.Save(user)
	if err != nil {
		return domain.User{}, fmt.Errorf("%s: %w", ErrorSave, err)
	}
	return user, nil
}

func (a authService) Login(user requests.LoginAuth) error {
	u, err := a.userService.FindByEmail(user.Email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return fmt.Errorf("%s: %w", ErrorLoginUserNotExist, err)
		}
		return fmt.Errorf("%s: %w", ErrorLoginInvalid, err)
	}
	valid := a.checkPasswordHash(user.Password, u.Password)
	if !valid {
		return fmt.Errorf("%s: %w", ErrorLoginInvalid, err)
	}
	return nil
}

func (a authService) checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
