package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

const (
	TEST  = "test"
	LOCAL = "local"
	DEV   = "development"
	PROD  = "production"
)

type configSchema struct {
	// Environment
	Environment string

	// Server configs
	Port string `required:"true"`

	// Database Configs (Postgres)
	DatabaseDSN      string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

var GtConfig *configSchema

func Configure() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var c configSchema
	err = envconfig.Process("GT", &c)
	if err != nil {
		log.Fatalf("failed to parse env file - err: %v", err)
	}

	GtConfig = &c
}
