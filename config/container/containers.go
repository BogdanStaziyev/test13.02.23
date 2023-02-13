package container

import (
	"context"
	"github.com/test_crud/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Container struct {
	Services
	Handlers
	Middleware
}

type Services struct {
}

type Handlers struct {
}

type Middleware struct {
}

func New(conf config.Configuration) Container {
	//coll := getDb(conf)

	return Container{
		Services:   Services{},
		Handlers:   Handlers{},
		Middleware: Middleware{},
	}
}

func getDb(conf config.Configuration) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf.MongoURL))
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("testProject")
	return coll
}
