package handler

import (
	"server/internal/crypto"
	"server/internal/repository"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	postgres *repository.Postgres
	aes      *crypto.Cryptor
)

func Set(db *repository.Postgres, cryptor *crypto.Cryptor) {
	postgres = db
	aes = cryptor
}

func Ping(ctx *gin.Context) {
	ctx.JSON(200, ResponseOk(
		"pong :)",
		ctx.ClientIP(),
	))
}

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
