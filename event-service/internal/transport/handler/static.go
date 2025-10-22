package handler

import (
	"server/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type (
	Response struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Payload *Payload `json:"payload"`
	}
	Payload struct {
		Timestamp int64  `json:"timestamp"`
		Client    string `json:"client"`
	}
)

var (
	postgres *repository.Postgres
)

func Set(db *repository.Postgres) {
	postgres = db
}

func Ping(ctx *gin.Context) {
	ctx.JSON(200, ResponseOk(
		"pong :)",
		ctx.ClientIP(),
	))
}

func ResponseOk(message string, client string) *Response {
	return &Response{
		Status:  "ok",
		Message: message,
		Payload: &Payload{
			Timestamp: time.Now().Unix(),
			Client:    client,
		},
	}
}

func ResponseError(err string, client string) *Response {
	return &Response{
		Status:  "error",
		Message: err,
		Payload: &Payload{
			Timestamp: time.Now().Unix(),
			Client:    client,
		},
	}
}
