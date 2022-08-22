package repository

import (
	"context"
	"fmt"
	"github.com/mrbelka12000/artforintrovert_testEx/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type manager struct {
	client *mongo.Client
}

func newManager(client *mongo.Client) *manager {
	return &manager{client}
}

func (m *manager) GetAll() ([]models.Product, error) {
	tempData := GetData(m.client)
	if data == nil {
		return nil, fmt.Errorf("failed to get data: %w", ErrNoData)
	}

	return tempData, nil
}

func (m *manager) Change(product *models.Product) (*models.Product, error) {
	return nil, nil
}

func (m *manager) Delete(id int) error {
	return nil
}

//Insert inserts default values to database.
func (m *manager) Insert() error {
	cfg := config.GetConf()

	coll := m.client.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection)
	var products = []models.Product{
		{
			ID:    primitive.NewObjectID(),
			Name:  "apple",
			Price: 12,
		},
		{
			ID:    primitive.NewObjectID(),
			Name:  "carrot",
			Price: 25,
		},
		{
			ID:    primitive.NewObjectID(),
			Name:  "milk",
			Price: 100,
		},
	}
	for _, v := range products {
		log.Println(coll.InsertOne(context.TODO(), v))
	}
	return nil
}
