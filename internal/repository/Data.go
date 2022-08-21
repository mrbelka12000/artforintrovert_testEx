package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

var (
	data    []models.Product
	updated bool = true
)

func updateData(client *mongo.Client) {
	// data = GetAll

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
