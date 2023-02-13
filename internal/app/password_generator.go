package app

import "golang.org/x/crypto/bcrypt"

type Generator interface {
	GeneratePasswordHash(password string) (string, error)
}

type generatePasswordHash struct {
	cost int
}

func NewGeneratePasswordHash(cost int) Generator {
	return generatePasswordHash{
		cost: cost,
	}
}

func (g generatePasswordHash) GeneratePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), g.cost)
	return string(bytes), err
}
