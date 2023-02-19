package container

import (
	"context"
	"github.com/test_crud/config"
	"github.com/test_crud/internal/app"
	"github.com/test_crud/internal/infra/database"
	"github.com/test_crud/internal/infra/http/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type Container struct {
	Services
	Handlers
}

type Services struct {
	app.UserService
	app.AuthService
}

type Handlers struct {
	handlers.AuthHandler
	handlers.UserHandler
}

func New(conf config.Configuration) Container {
	coll := getDb(conf)

	userRepository := database.NewUSerRepo(*coll)
	passwordGenerator := app.NewGeneratePasswordHash(bcrypt.DefaultCost)
	userService := app.NewUserService(userRepository, passwordGenerator)
	authService := app.NewAuthService(userService, conf)
	authController := handlers.NewAuthHandler(authService)
	userController := handlers.NewUserHandler(userService)

	return Container{
		Services: Services{
			userService,
			authService,
		},
		Handlers: Handlers{
			authController,
			userController,
		},
	}
}

func getDb(conf config.Configuration) *mongo.Database {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.MongoURL))
	if err != nil {
		log.Println(err)
	}
	coll := client.Database(conf.DbName)
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("Could not connect to mongo db %s", err)
	}
	return coll
}
