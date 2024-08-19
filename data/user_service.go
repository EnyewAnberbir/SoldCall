package data

import (
	"context"
	"fmt"
	"log"
    "os"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	client     *mongo.Client
	Collection *mongo.Collection
)

func InitMongoDB() error {
	//var err error
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    // Get the MongoDB URI from the environment variables
    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatalf("MONGODB_URI not set in environment variables")
    }

    // Set client options and connect to the MongoDB server
    clientOptions := options.Client().ApplyURI(mongoURI)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	// Check the connection
	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	Collection = client.Database("testdb").Collection("users")
	return nil
}
