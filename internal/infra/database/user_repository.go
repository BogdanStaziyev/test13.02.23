package database

import (
	"context"
	"fmt"
	"github.com/test_crud/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
	"time"
)

const (
	UsersTable       = "users"
	ErrorSave        = "user repository save user"
	ErrorFindByEmail = "user repository find by email user"
	ErrorFindAll     = "user repository find all users"
)

var ctx = context.TODO()

type userToDomain struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Email       string             `bson:"email"`
	Name        string             `bson:"name"`
	Password    string             `bson:"password,omitempty"`
	CreatedDate time.Time          `bson:"created_date,omitempty"`
	UpdatedDate time.Time          `bson:"updated_date"`
}

type UserRepo interface {
	Save(user domain.User) error
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

func (u userRepo) Save(user domain.User) error {
	domainUser := u.mapDomainToModel(user)
	domainUser.CreatedDate = time.Now()
	domainUser.UpdatedDate = time.Now()
	_, err := u.coll.InsertOne(ctx, &domainUser)
	if err != nil {
		return fmt.Errorf("%s: %w", ErrorSave, err)
	}
	return nil
}

func (u userRepo) FindByEmail(email string) (domain.User, error) {
	var domainUser userToDomain
	email = strings.ToLower(email)
	err := u.coll.FindOne(ctx, bson.M{"email": email}).Decode(&domainUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("%s: %w", ErrorFindByEmail, err)
	}
	return u.mapModelToDomain(domainUser), nil
}

func (u userRepo) FindAll() ([]domain.User, error) {
	var users []userToDomain
	find, err := u.coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", ErrorFindAll, err)
	}
	for find.Next(ctx) {
		var user userToDomain
		err = find.Decode(&user)
		if err != nil {
			return []domain.User{}, fmt.Errorf("%s: %w", ErrorFindAll, err)
		}
		users = append(users, user)
	}
	err = find.Close(ctx)
	if err != nil {
		log.Println(err)
	}
	return u.mapUsersCollection(users), nil
}

func (u userRepo) mapDomainToModel(d domain.User) userToDomain {
	return userToDomain{
		ID: func(s string) primitive.ObjectID {
			id, err := primitive.ObjectIDFromHex(s)
			if err != nil {
				log.Println(err)
			}
			return id
		}(d.ID),
		Email:    strings.ToLower(d.Email),
		Password: d.Password,
		Name:     d.Name,
	}
}

func (u userRepo) mapModelToDomain(d userToDomain) domain.User {
	return domain.User{
		ID:          d.ID.Hex(),
		Email:       d.Email,
		Password:    d.Password,
		Name:        d.Name,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
	}
}

func (u userRepo) mapUsersCollection(users []userToDomain) []domain.User {
	var result []domain.User
	for _, coll := range users {
		newUser := u.mapModelToDomain(coll)
		result = append(result, newUser)
	}
	return result
}
