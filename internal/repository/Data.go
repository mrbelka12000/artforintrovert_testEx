package repository

import (
	"context"
	"github.com/mrbelka12000/artforintrovert_testEx/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

var (
	data    []models.Product
	updated = true
)

func updateData(client *mongo.Client) {
	cfg := config.GetConf()

	coll := client.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection)

	var products []models.Product

	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		zap.S().Errorf("unable to find data: %v", err)
		return
	}

	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			zap.S().Debugf("decode error: %v", err)
			continue
		}

		products = append(products, product)
	}

	err = cursor.Err()
	if err != nil {
		zap.S().Debugf("cursor error: %v", err)
		return
	}

	data = products

	updated = false
}

func GetData(client *mongo.Client) []models.Product {
	if !updated {
		return data
	}

	updateData(client)

	if updated {
		zap.S().Debug("data can not be updated")
		return nil
	}

	zap.S().Info("data updated")
	return data
}
