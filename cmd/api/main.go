package main

import (
	"log"

	"github.com/jdejesus007/gt-api-project/api"
	"github.com/jdejesus007/gt-api-project/internal/config"
)

//	@title			Gt API
//	@version		0.0.1
//	@description	Gt Web API

//	@host		localhost:3000
//	@BasePath	/

func main() {
	log.Println("vim-go: GT World!")

	// Init config obj
	config.Configure()

	// Setup and run API rest endpoints
	api := api.NewBuilder().
		WithRepositoryProvider(api.GetRepositories()).
		Finalize()

	// Open and serve web server
	if err := api.ListenAndServe(); err != nil {
		log.Fatalln("ListenAndServe Failed! Fatal Error: ", err)
	}
}
