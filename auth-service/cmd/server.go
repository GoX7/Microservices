package main

import (
	"server/internal/config"
	"server/internal/repository"
	"server/internal/transport"
	"server/models"

	"github.com/gin-gonic/gin"
)

func main() {
	configm := new(models.Config)
	config.NewConfig(configm)

	postgres := new(repository.Postgres)
	repository.NewPostgres(configm, postgres)
	postgres.SetTable()

	engine := gin.Default()
	transport.Register(engine, configm, postgres)
	transport.Listen(engine, configm)
}
