package main

import (
	"fmt"
	handler "go-deck-of-cards/internal/app/handler/deck"
	"go-deck-of-cards/internal/app/middleware"
	"go-deck-of-cards/internal/pkg/config"
	"go-deck-of-cards/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitMongoDB()

	router := gin.Default()
	h := handler.NewDeckHandler()
	m := middleware.UuidMiddleware

	router.Use(m)
	h.RegisterRoutes(router)

	router.Run(fmt.Sprintf("%v:%s", config.ServerHost, config.ServerPort))
}
