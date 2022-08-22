package repository

import (
	"context"
	"github.com/mrbelka12000/artforintrovert_testEx/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

var (
	data         []models.Product
	needToUpdate = true
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

	needToUpdate = false
}

func GetData(client *mongo.Client) []models.Product {
	if !needToUpdate {
		return data
	}

	updateData(client)

	if needToUpdate {
		zap.S().Debug("data can not be needToUpdate")
		return nil
	}

	zap.S().Info("data needToUpdate")
	return data
}

func Updater(client *mongo.Client, ctx context.Context, ch chan struct{}) {
	ticker := time.NewTicker(5 * time.Second)

	for {

		select {
		case <-ticker.C:
			if needToUpdate {
				updateData(client)
			}
		case <-ctx.Done():
			zap.S().Info("updater func successfully ended")
			ch <- struct{}{}
			return
		}
	}

}
