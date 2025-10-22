package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type (
	RequestEvent struct {
		Event string `json:"event"`
	}
)

func PingEvent(ctx *gin.Context) {
	if cores.EventService.Ping() {
		ctx.JSON(200, ResponseOk(
			"event service is lived",
			ctx.ClientIP(),
		))
	} else {
		ctx.JSON(200, ResponseError(
			"event servcie is dead",
			ctx.ClientIP(),
		))
	}
}

func RegisterEvent(ctx *gin.Context) {
	login := ctx.GetString("login")
	if login == "*401" {
		ctx.JSON(401, ResponseError(
			"Aw, unauthorizathion!",
			ctx.ClientIP(),
		))
		return
	}

	var request RequestEvent
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(400, ResponseError(
			"Aw, json struct is incorect",
			ctx.ClientIP(),
		))
		return
	}

	fmt.Println("event", request.Event)
	id, err := cores.EventService.Register(login, request.Event)
	if err != nil {
		ctx.JSON(400, ResponseError(
			err.Error(),
			ctx.ClientIP(),
		))
		return
	}

	fmt.Println("id", id)
	ctx.JSON(201, ResponseOk(
		"create, id="+id,
		ctx.ClientIP(),
	))
}

func AllEvent(ctx *gin.Context) {
	login := ctx.GetString("login")
	if login == "*401" {
		ctx.JSON(401, ResponseError(
			"Aw, unauthorizathion!",
			ctx.ClientIP(),
		))
		return
	}

	events, err := cores.EventService.GetEvents(login)
	if err != nil {
		ctx.JSON(400, ResponseError(
			err.Error(),
			ctx.ClientIP(),
		))
		return
	}

	ctx.JSON(200, events)
}

func SearchEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, ResponseError(
			"Aw, none id, usage /event/[id]",
			ctx.ClientIP(),
		))
		return
	}

	login := ctx.GetString("login")
	if login == "*401" {
		ctx.JSON(401, ResponseError(
			"Aw, unauthorizathion!",
			ctx.ClientIP(),
		))
		return
	}

	event, err := cores.EventService.GetEvent(id, login)
	if err != nil {
		ctx.JSON(400, ResponseError(
			err.Error(),
			ctx.ClientIP(),
		))
		return
	}

	ctx.JSON(200, event)
}

func DeleteEvent(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(400, ResponseError(
			"Aw, none id, usage /event/[id]",
			ctx.ClientIP(),
		))
		return
	}

	login := ctx.GetString("login")
	if login == "*401" {
		ctx.JSON(401, ResponseError(
			"Aw, unauthorizathion!",
			ctx.ClientIP(),
		))
		return
	}

	if err := cores.EventService.DeleteEvent(id, login); err != nil {
		ctx.JSON(400, ResponseError(
			err.Error(),
			ctx.ClientIP(),
		))
	}

	ctx.JSON(200, ResponseOk(
		"Delete",
		ctx.ClientIP(),
	))
}
