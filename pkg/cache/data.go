// Package cache save and pdates the data as needed with data from database when server starting.
package cache

import (
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
	"github.com/mrbelka12000/artforintrovert_testEx/internal/models"
)

type data struct {
	data         []models.Product
	needToUpdate bool
	sync.Mutex
}

var globData = &data{
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
		log.Println("error: data can not be updated")
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
			log.Println("updater func successfully ended")
			ch <- struct{}{}
			return
		}
	}
}

func updateData(client *mongo.Client) {
	cfg, _ := config.GetConf()

	coll := client.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection)

	var products []models.Product

	cursor, err := coll.Find(context.Background(), bson.D{})
	if err != nil {
		log.Printf("unable to find data: %v \n", err)
		return
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var product models.Product
		err := cursor.Decode(&product)
		if err != nil {
			log.Printf("decode error: %v \n", err)
			continue
		}

		products = append(products, product)
	}

	err = cursor.Err()
	if err != nil {
		log.Printf("cursor error: %v \n", err)
		return
	}

	log.Println("data successfully updated")

	globData.data = products
	globData.needToUpdate = false
}
