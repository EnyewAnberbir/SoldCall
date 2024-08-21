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
	client          *mongo.Client
	UserCollection  *mongo.Collection
	EmojiCollection *mongo.Collection
	ContactCollection *mongo.Collection
	BusinessCollection *mongo.Collection
)

func InitMongoDB() error {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    mongoURI := os.Getenv("MONGODB_URI")
    if mongoURI == "" {
        log.Fatalf("MONGODB_URI not set in environment variables")
    }

    clientOptions := options.Client().ApplyURI(mongoURI)
	client, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return err
	}

	fmt.Println("Connected to MongoDB!")
	UserCollection = client.Database("testdb").Collection("users")
	EmojiCollection = client.Database("testdb").Collection("emojis")
	ContactCollection = client.Database("testdb").Collection("contacts")
	BusinessCollection = client.Database("testdb").Collection("businesses")
	
	return nil
}
