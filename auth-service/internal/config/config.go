package config

import (
	"os"
	"server/models"
)

func NewConfig(config *models.Config) {
	*config = models.Config{
		SERVER_ADDR:       os.Getenv("SERVER_ADDR"),
		SERVER_CRYPTO:     os.Getenv("SERVER_CRYPTO"),
		POSTGRES_HOST:     os.Getenv("POSTGRES_HOST"),
		POSTGRES_USER:     os.Getenv("POSTGRES_USER"),
		POSTGRES_PASSWORD: os.Getenv("POSTGRES_PASSWORD"),
		POSTGRES_DBNAME:   os.Getenv("POSTGRES_DBNAME"),
	}
}
