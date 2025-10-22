package config

import (
	"os"
	"server/models"
)

func NewConfig(config *models.Config) {
	*config = models.Config{
		SERVER_ADDR:   os.Getenv("SERVER_ADDR"),
		SERVICE_AUTH:  os.Getenv("SERVICE_AUTH"),
		SERVICE_EVENT: os.Getenv("SERVICE_EVENT"),
	}
}
