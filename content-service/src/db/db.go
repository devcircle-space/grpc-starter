package db

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance      *mongo.Client
	clientInstanceError error
	mongoOnce           sync.Once
)

const (
	DATABASE_URL          = "DATABASE_URL"
	MONGO_CONNECT_TIMEOUT = 10 * time.Second
)

// GetMongoClient returns a singleton MongoDB client
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		uri := os.Getenv(DATABASE_URL)

		ctx, cancel := context.WithTimeout(context.Background(), MONGO_CONNECT_TIMEOUT)
		defer cancel()
		clientOptions := options.Client().ApplyURI(uri)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			clientInstanceError = err
			return
		}
		// Ping to verify connection
		if err := client.Ping(ctx, nil); err != nil {
			clientInstanceError = err
			return
		}
		log.Println("Connected to MongoDB at", uri)
		clientInstance = client
	})
	return clientInstance, clientInstanceError
}
