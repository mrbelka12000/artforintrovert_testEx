package mongodb

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type Data struct {
	data         []models.Product
	needToUpdate bool
	sync.Mutex
}

var globData = &Data{
	needToUpdate: true,
}

const waitLimitForTicker = 1 * time.Minute

func GetData(client *mongo.Client) []models.Product {
	globData.Lock()
	defer globData.Unlock()
	if !globData.needToUpdate {
		return globData.data
	}
	updateData(client)

	if globData.needToUpdate {
		zap.S().Warn("data can not be updated")
		return nil
	}

	return globData.data
}

func NeetToUpdate() {
	globData.Lock()
	globData.needToUpdate = true
	globData.Unlock()
}

func Updater(client *mongo.Client, ctx context.Context, ch chan struct{}) {
	ticker := time.NewTicker(waitLimitForTicker)

	globData.Lock()
	updateData(client)
	globData.Unlock()

	for {
		select {
		case <-ticker.C:
			globData.Lock()
			if globData.needToUpdate {
				updateData(client)
			}
			globData.Unlock()
		case <-ctx.Done():
			zap.S().Info("updater func successfully ended")
			ch <- struct{}{}
			return
		}
	}
}

func updateData(client *mongo.Client) {
	cfg := config.GetConf()

	coll := client.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection)

	var products []models.Product

	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		zap.S().Errorf("unable to find data: %v", err)
		return
	}
	defer cursor.Close(context.Background())

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

	zap.S().Info("data successfully updated")

	globData.data = products
	globData.needToUpdate = false
}
