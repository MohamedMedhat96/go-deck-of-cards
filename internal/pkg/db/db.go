package db

import (
	"context"
	"log"
	"time"

	"go-deck-of-cards/internal/pkg/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongoDB() {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(config.MongoDbURI)

	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}

func GetMongoDB() *mongo.Database {
	return client.Database(config.DatabaseName)
}

func GetMongoClient() *mongo.Client {
	return client
}

func CloseConnection(ctx *gin.Context) {
	client.Disconnect(ctx)
}

func GetCollection(collectionName string) *mongo.Collection {
	return GetMongoDB().Collection(collectionName)
}
