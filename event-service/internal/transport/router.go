package transport

import (
	"fmt"
	"net/http"
	"server/internal/repository"
	"server/internal/transport/handler"
	"server/models"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine, db *repository.Postgres) {
	engine.GET("/ping", handler.Ping)

	handler.Set(db)
	engine.POST("/register", handler.RegisterEvent)
	engine.GET("/events/:login", handler.GetEvents)
	engine.GET("/events/:login/:id", handler.GetEvent)
	engine.DELETE("/events/:login/:id", handler.DeleteEvent)
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
