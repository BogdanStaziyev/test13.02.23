package database

import (
	"context"
	"fmt"
	"github.com/test_crud/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"strings"
	"time"
)

const UsersTable = "users"

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
	//FindByID(id int64) (domain.User, error)
	//Delete(id int64) error
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
		return domain.User{}, err
	}
	log.Println(res.InsertedID)
	return u.mapModelToDomain(domainUser), nil
}

func (u userRepo) FindByEmail(email string) (domain.User, error) {
	var domainUser user
	email = strings.ToLower(email)
	err := u.coll.FindOne(context.Background(), bson.M{"email": email}).Decode(&domainUser)
	if err != nil {
		return domain.User{}, fmt.Errorf("user repository save user: %w", err)
	}
	return u.mapModelToDomain(domainUser), nil
}

//func (u userRepo) FindByID(id string) (domain.User, error) {
//	var domainUser user
//
//	err := u.coll.FindOne(context.Background(), bson.M{"_id": id})
//	if err != nil {
//		return domain.User{}, fmt.Errorf("user repository finde by id user: %w", err)
//	}
//	return u.mapModelToDomain(domainUser), nil
//}
//
//func (u userRepo) Delete(id int64) error {
//	err := u.coll.Find(db.Cond{
//		"id":           id,
//		"deleted_date": nil,
//	}).Update(map[string]interface{}{"deleted_date": time.Now()})
//	if err != nil {
//		return fmt.Errorf("user repository delete user: %w", err)
//	}
//	return nil
//}

func (u userRepo) FindAll() ([]domain.User, error) {
	var users []user
	find, err := u.coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for find.Next(context.Background()) {
		var us user
		err = find.Decode(&us)
		if err != nil {
			return []domain.User{}, err
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
