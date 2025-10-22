package handler

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type (
	RequestRegisterAuth struct {
		Login    string `json:"login" validate:"required,min=3,max=24,alphanum"`
		Password string `json:"password" validate:"required"`
	}
)

func RegisterAuth(ctx *gin.Context) {
	request := new(RequestRegisterAuth)
	if err := ctx.BindJSON(request); err != nil {
		ctx.String(400, "Aw, json struct is incorect")
		return
	}

	if err := validator.New().Struct(request); err != nil {
		ctx.String(400, "Aw, incorect symbol(s)")
		return
	}

	hashed := sha256.Sum256([]byte(request.Password))
	request.Password = base64.StdEncoding.EncodeToString(hashed[:])

	lastId, err := postgres.RegisterUser(request.Login, request.Password)
	if err != nil {
		ctx.String(400, err.Error())
		return
	}

	cipherText := fmt.Sprintf("%d|%s|%s", lastId, request.Login, request.Password)
	str := aes.Seal([]byte(cipherText))
	ctx.String(201, str)
}

func LoginAuth(ctx *gin.Context) {
	request := new(RequestRegisterAuth)
	if err := ctx.BindJSON(request); err != nil {
		ctx.String(400, "Aw, json struct is incorect")
		return
	}

	if err := validator.New().Struct(request); err != nil {
		ctx.String(400, "Aw, incorect symbol(s)")
		return
	}

	hashed := sha256.Sum256([]byte(request.Password))
	request.Password = base64.StdEncoding.EncodeToString(hashed[:])

	lastId, err := postgres.LoginUser(request.Login, request.Password)
	if err != nil {
		ctx.String(400, err.Error())
		return
	}

	cipherText := fmt.Sprintf("%d|%s|%s", lastId, request.Login, request.Password)
	str := aes.Seal([]byte(cipherText))
	ctx.String(200, str)
}

func TokenAuth(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.String(500, "Aw, server.io error, contact the admin")
		return
	}

	user, err := aes.Open(string(body))
	if err != nil {
		ctx.String(400, "Aw, token is incorect")
		return
	}

	datas := strings.Split(string(user), "|")
	id, err := postgres.LoginUser(datas[1], datas[2])
	if err != nil {
		ctx.String(400, err.Error())
		return
	}

	if fmt.Sprint(id) != datas[0] {
		ctx.String(400, "Aw, account is outdated")
		return
	}

	ctx.JSON(200, gin.H{
		"id":    datas[0],
		"login": datas[1],
	})
}
