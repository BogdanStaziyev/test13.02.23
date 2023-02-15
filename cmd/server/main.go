package main

import (
	"github.com/test_crud/config"
	"github.com/test_crud/config/container"
	"github.com/test_crud/internal/infra/http"
	"log"
)

func main() {
	var conf = config.GetConfiguration()

	// Echo Server
	srv := http.NewServer()

	cont := container.New(conf)

	http.EchoRouter(srv.Echo, cont)

	err := srv.Start(conf.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
