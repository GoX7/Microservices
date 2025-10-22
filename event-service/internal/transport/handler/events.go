package handler

import (
	"fmt"
	"server/internal/repository"

	"github.com/gin-gonic/gin"
)

type (
	RegisterEventRequest struct {
		Login string `json:"login"`
		Event string `json:"event"`
	}
	UserEventRequest struct {
		Id    int64  `json:"id"`
		Login string `json:"login"`
	}

	EventsResponse struct {
		Events []*repository.Event `json:"events"`
	}
)

func RegisterEvent(ctx *gin.Context) {
	var request RegisterEventRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.String(400, "Aw, json request is incorect")
		return
	}

	id, err := postgres.CreateEvent(
		request.Login,
		request.Event,
	)
	if err != nil {
		ctx.String(400, err.Error())
		return
	}

	ctx.String(201, fmt.Sprint(id))
}

func GetEvents(ctx *gin.Context) {
	login := ctx.Param("login")
	if login == "" {
		ctx.String(400, "login")
		return
	}

	events, err := postgres.GetEvents(login)
	if err != nil {
		ctx.String(400, err.Error())
		return
	}

	ctx.JSON(200, EventsResponse{Events: events})
}

func GetEvent(ctx *gin.Context) {
	login := ctx.Param("login")
	if login == "" {
		ctx.String(400, "login")
		return
	}

	id := ctx.Param("id")
	if id == "" {
		ctx.String(400, "id")
		return
	}

	event, err := postgres.GetEvent(id, login)
	if err != nil {
		ctx.String(400, err.Error())
		return
	}

	ctx.JSON(200, event)
}

func DeleteEvent(ctx *gin.Context) {
	login := ctx.Param("login")
	if login == "" {
		ctx.String(400, "login")
		return
	}

	id := ctx.Param("id")
	if id == "" {
		ctx.String(400, "id")
		return
	}

	if err := postgres.DeleteEvent(id, login); err != nil {
		ctx.String(400, err.Error())
		return
	}

	ctx.Status(200)
}
