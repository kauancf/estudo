package main

import (
	"log"

	"github.com/kauancf/estudo/tree/main/api_students/api"
)

func main() {

	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
