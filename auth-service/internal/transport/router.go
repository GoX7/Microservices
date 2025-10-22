package transport

import (
	"fmt"
	"net/http"
	"server/internal/crypto"
	"server/internal/repository"
	"server/internal/transport/handler"
	"server/models"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine, config *models.Config, postgres *repository.Postgres) {
	handler.Set(
		postgres,
		crypto.NewCryptor([]byte(config.SERVER_CRYPTO)),
	)

	engine.GET("/ping", handler.Ping)
	engine.POST("/register", handler.RegisterAuth)
	engine.POST("/login", handler.LoginAuth)
	engine.POST("/token", handler.TokenAuth)
}

func Listen(engine *gin.Engine, config *models.Config) {
	server := http.Server{
		Handler: engine,
		Addr:    config.SERVER_ADDR,
	}

	fmt.Println("[+] server.listen...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println("[-] server.listen: " + err.Error())
	}
}
