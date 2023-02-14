package database

import (
	"context"
	"fmt"
	"github.com/test_crud/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
	"time"
)

const (
	UsersTable       = "users"
	ErrorSave        = "user repository save user"
	ErrorFindByEmail = "user repository find by email user"
	ErrorFindAll     = "user repository find all users"
)

type user struct {
	ID          string     `bson:"_id,omitempty"`
	Email       string     `bson:"email"`
	Name        string     `bson:"name"`
	Password    string     `bson:"password,omitempty"`
	CreatedDate time.Time  `bson:"created_date,omitempty"`
	UpdatedDate time.Time  `bson:"updated_date"`
	DeletedDate *time.Time `bson:"deleted_date,omitempty"`
}

type UserRepo interface {
	Save(user domain.User) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
	FindAll() ([]domain.User, error)
}

type userRepo struct {
	coll mongo.Collection
}

func NewUSerRepo(db mongo.Database) UserRepo {
	return userRepo{
		coll: *db.Collection(UsersTable),
	}
}

func (u userRepo) Save(user domain.User) (domain.User, error) {
	domainUser := u.mapDomainToModel(user)
	domainUser.CreatedDate = time.Now()
	domainUser.UpdatedDate = time.Now()
	res, err := u.coll.InsertOne(context.Background(), &domainUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("%s: %w", ErrorSave, err)
	}
	domainUser.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return u.mapModelToDomain(domainUser), nil
}

func (u userRepo) FindByEmail(email string) (domain.User, error) {
	var domainUser user
	email = strings.ToLower(email)
	err := u.coll.FindOne(context.Background(), bson.M{"email": email}).Decode(&domainUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("%s: %w", ErrorFindByEmail, err)
	}
	return u.mapModelToDomain(domainUser), nil
}

func (u userRepo) FindAll() ([]domain.User, error) {
	var users []user
	find, err := u.coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrorFindAll, err)
	}
	for find.Next(context.Background()) {
		var us user
		err = find.Decode(&us)
		if err != nil {
			return []domain.User{}, fmt.Errorf("%s: %w", ErrorFindAll, err)
		}
		users = append(users, us)
	}
	return u.mapUsersCollection(users), nil
}

func (u userRepo) mapDomainToModel(d domain.User) user {
	return user{
		ID:       d.ID,
		Email:    strings.ToLower(d.Email),
		Password: d.Password,
		Name:     d.Name,
	}
}

func (u userRepo) mapModelToDomain(d user) domain.User {
	return domain.User{
		ID:          d.ID,
		Email:       d.Email,
		Password:    d.Password,
		Name:        d.Name,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (u userRepo) mapUsersCollection(users []user) []domain.User {
	var result []domain.User
	for _, coll := range users {
		newUser := u.mapModelToDomain(coll)
		result = append(result, newUser)
	}
	return result
}
