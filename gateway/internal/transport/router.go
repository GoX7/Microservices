package transport

import (
	"fmt"
	"net/http"
	"server/internal/core"
	"server/internal/transport/handler"
	"server/internal/transport/middleware"
	"server/models"

	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine, core *core.Core) {
	handler.LoadResource(core)

	engine.Use(middleware.NewAuth(core))

	engine.GET("/auth/ping", handler.PingAuth)
	engine.GET("/event/ping", handler.PingEvent)

	engine.POST("/auth/register", handler.RegisterAuth)
	engine.POST("/auth/login", handler.LoginAuth)
	engine.GET("/auth/me", handler.AccountAuth)

	engine.POST("/events/register", handler.RegisterEvent)
	engine.DELETE("/events/:id", handler.DeleteEvent)
	engine.GET("/events/:id", handler.SearchEvent)
	engine.GET("/events", handler.AllEvent)
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
