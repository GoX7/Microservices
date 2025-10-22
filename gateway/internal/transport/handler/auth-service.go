package handler

import (
	"github.com/gin-gonic/gin"
)

type (
	RequestAuth struct {
		Login    string `json:"login" validate:"required,alphanum,min=3,max=24"`
		Password string `json:"password" validate:"required"`
	}
)

func PingAuth(ctx *gin.Context) {
	if cores.AuthService.Ping() {
		ctx.JSON(200, ResponseOk(
			"auth service is lived",
			ctx.ClientIP(),
		))
	} else {
		ctx.JSON(200, ResponseError(
			"auth servcie is dead",
			ctx.ClientIP(),
		))
	}
}

func RegisterAuth(ctx *gin.Context) {
	request := new(RequestAuth)
	if err := ctx.BindJSON(request); err != nil {
		ctx.JSON(400, ResponseError(
			`request is not valide struct, usage: {"login": [login], "password": [password]}`,
			ctx.ClientIP(),
		))
		return
	}

	userToken, err := cores.AuthService.Register(
		request.Login,
		request.Password,
	)
	if err != nil {
		ctx.JSON(400, ResponseError(
			err.Error(),
			ctx.ClientIP(),
		))
		return
	}

	ctx.JSON(201, ResponseOk(
		userToken,
		ctx.ClientIP(),
	))
}

func LoginAuth(ctx *gin.Context) {
	request := new(RequestAuth)
	if err := ctx.BindJSON(request); err != nil {
		ctx.JSON(400, ResponseError(
			`request is not valide struct, usage: {"login": [login], "password": [password]}`,
			ctx.ClientIP(),
		))
		return
	}

	userToken, err := cores.AuthService.Login(
		request.Login,
		request.Password,
	)
	if err != nil {
		ctx.JSON(400, ResponseError(
			err.Error(),
			ctx.ClientIP(),
		))
	}

	ctx.JSON(200, ResponseOk(
		userToken,
		ctx.ClientIP(),
	))
}

func AccountAuth(ctx *gin.Context) {
	id := ctx.GetString("id")
	login := ctx.GetString("login")

	if id == "*401" && login == "*401" {
		ctx.JSON(401, ResponseError(
			"Aw, Unauthorizathion!",
			ctx.ClientIP(),
		))
	} else {
		ctx.JSON(200, ResponseAccount(
			id, login,
			ctx.ClientIP(),
		))
	}
}
