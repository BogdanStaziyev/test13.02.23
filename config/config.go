package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
)

type Configuration struct {
	MongoURL string
}

func GetConfiguration() Configuration {

	err := godotenv.Load(filepath.Join(".env"))
	if err != nil {
		log.Println(err)
	}

	return Configuration{
		MongoURL: os.Getenv("MONGOURL"),
	}
}
