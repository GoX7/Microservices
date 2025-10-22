package middleware

import (
	"server/internal/core"
	"strings"

	"github.com/gin-gonic/gin"
)

func NewAuth(cores *core.Core) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("id", "*401")
		ctx.Set("login", "*401")

		header := ctx.GetHeader("Authorization")
		if header != "" {
			auth := strings.Split(header, " ")
			if auth[0] == "Bearer" {
				if auth[1] != "" {
					user, err := cores.AuthService.LoginToken(auth[1])
					if err == nil {
						ctx.Set("id", user.Id)
						ctx.Set("login", user.Login)
					}
				}
			}
		}

		ctx.Next()
	}
}
