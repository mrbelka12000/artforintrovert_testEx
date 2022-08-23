package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
)

const waitLimit = 8 * time.Second

func GetMongoDBClient() (*mongo.Client, error) {
	cfg := config.GetConf()

	ctx, _ := context.WithTimeout(context.Background(), waitLimit)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoDB.MongoUrl))
	if err != nil {
		return nil, fmt.Errorf("can not connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	return client, err
}

