package core

import (
	"fmt"
	"net/http"
	"server/models"
)

type (
	Core struct {
		AuthService  *AuthService
		EventService *EventService
	}
	AuthService struct {
		Addr string
	}
	EventService struct {
		Addr string
	}
)

func pingService(addr string) bool {
	response, err := http.Get("http://" + addr + "/ping")
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return false
	}

	if response.StatusCode != 200 {
		return false
	} else {
		return true
	}
}

func NewCore(config *models.Config, core *Core) {
	auth := pingService(config.SERVICE_AUTH)
	event := pingService(config.SERVICE_EVENT)

	if auth {
		fmt.Println("[+] core.auth: ping")
	}
	if event {
		fmt.Println("[+] core.event: ping")
	}

	*core = Core{
		AuthService: &AuthService{
			Addr: config.SERVICE_AUTH,
		},
		EventService: &EventService{
			Addr: config.SERVICE_EVENT,
		},
	}
}
