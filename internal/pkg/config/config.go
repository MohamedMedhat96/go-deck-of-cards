package config

import (
	"os"
)

var (
	ServerPort string

	MongoDbURI   string
	DatabaseName string
)

func init() {
	ServerPort = getEnv("SERVER_PORT", "8080")

	MongoDbURI = getEnv("MONGO_DB_URI", "mongodb://localhost:27017/")

	DatabaseName = getEnv("DATABASE_NAME", "deck")
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	return value
}
