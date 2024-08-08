package main

import (
	"github.com/rs/zerolog/log"

	"github.com/kauancf/estudo/tree/main/api_students/api"
)

func main() {

	server := api.NewServer()

	server.ConfigureRoutes()

	if err := server.Start(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")

	}
}
