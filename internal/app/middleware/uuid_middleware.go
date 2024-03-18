package middleware

import (
	"github.com/gin-gonic/gin"
)

func UuidMiddleware(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	if uuid != "" {
		ctx.Set("uuid", uuid)
	}
}
