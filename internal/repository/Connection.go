package repository

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const waitLimit = 8 * time.Second

func GetMongoDBClient() (*mongo.Client, error) {
	ctx, _ := context.WithTimeout(context.Background(), waitLimit)
	uri := os.Getenv("MONGODB_URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("can not connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	return client, err
}
