package config

import (
	"os"
)

type Configuration struct {
	MongoURL    string
	RedirectUrl string
	Port        string
	Host        string
	DbName      string
}

func GetConfiguration() Configuration {
	//
	//err := godotenv.Load(filepath.Join(".env"))
	//if err != nil {
	//	log.Println(err)
	//}

	return Configuration{
		MongoURL:    os.Getenv("MONGO_URL"),
		RedirectUrl: os.Getenv("REDIRECT_URL"),
		Port:        os.Getenv("PORT"),
		Host:        os.Getenv("HOST"),
		DbName:      os.Getenv("DB_NAME"),
	}
}
