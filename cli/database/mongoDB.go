package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectWithMongodb() *mongo.Client {
	if err := godotenv.Load("../.env"); err != nil {
		log.Fatal(err)
	}

	url := os.Getenv("MONGO_URL")
	if url == "" {
		log.Fatal("MONGO_URL is not set")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to mongodb")
	}

	return client
}

var Client *mongo.Client = ConnectWithMongodb()
