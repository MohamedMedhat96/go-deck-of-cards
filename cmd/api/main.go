package main

import (
	"go-deck-of-cards/internal/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitMongoDB()

	router := gin.Default()

	router.Run("localhost:8080")
}
