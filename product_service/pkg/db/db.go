package db

import (
	"EXAM3_with_mongodb/product_service/config"
	"context"
	"log"

	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Mclient *mongo.Client
}

func New(cfg config.Config) (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI("mongodb://mongodb:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database("users_products")

	return db, nil
}
