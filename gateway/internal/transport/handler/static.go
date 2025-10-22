package handler

import (
	"server/internal/core"
	"time"
)

var (
	cores *core.Core
)

func LoadResource(corem *core.Core) {
	cores = corem
}

type (
	User struct {
		Status  string   `json:"status"`
		Id      string   `json:"id"`
		Login   string   `json:"login"`
		Payload *Payload `json:"payload"`
	}
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
		Status:  "success",
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

func ResponseAccount(id string, login string, client string) *User {
	return &User{
		Status: "success",
		Login:  login, Id: id,
		Payload: &Payload{
			Timestamp: time.Now().Unix(),
			Client:    client,
		},
	}
}
