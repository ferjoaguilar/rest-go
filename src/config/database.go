package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongodbInstance struct {
	Client *mongo.Client
	DB *mongo.Database 
}

var Query MongodbInstance

func Connection() {
	// Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB"))
	// Connect mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	// check connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connected")

	Query = MongodbInstance{
		Client: client,
		DB: client.Database(os.Getenv("DATABASE")),
	}
}