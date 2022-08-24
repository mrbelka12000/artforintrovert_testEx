// Package mongodb implements connection to MongoDB.
package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
)

const waitLimitForConnection = 8 * time.Second

func GetMongoDBClient(ctx context.Context) (*mongo.Client, error) {
	cfg, _ := config.GetConf()

	connCtx, _ := context.WithTimeout(ctx, waitLimitForConnection)

	client, err := mongo.Connect(connCtx, options.Client().ApplyURI(cfg.MongoDB.MongoUrl))
	if err != nil {
		return nil, fmt.Errorf("can not connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("connection error: %w", err)
	}

	return client, err
}
