package db

import (
	"context"
	"log"

	"go-deck-of-cards/internal/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func InitMongoDB() {
	var err error

	clientOptions := options.Client().ApplyURI(config.MongoDbURI)

	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")
}

func GetMongoClient() *mongo.Client {
	return client
}
