package main

import (
	"server/internal/config"
	"server/internal/core"
	"server/internal/transport"
	"server/models"

	"github.com/gin-gonic/gin"
)

func main() {
	configm := new(models.Config)
	config.NewConfig(configm)

	corem := new(core.Core)
	core.NewCore(configm, corem)

	engine := gin.Default()
	transport.Register(engine, corem)
	transport.Listen(engine, configm)
}
