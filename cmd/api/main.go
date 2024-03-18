package main

import (
	handler "go-deck-of-cards/internal/app/handler/deck"
	"go-deck-of-cards/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitMongoDB()

	router := gin.Default()
	h := handler.NewDeckHandler()

	h.RegisterRoutesAndMiddleware(router)

	router.Run("localhost:8080")
}
