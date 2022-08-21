package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type manager struct {
	client *mongo.Client
}

func newManager(client *mongo.Client) *manager {
	return &manager{client}
}

func (m *manager) GetAll() ([]models.Product, error) {
	return nil, nil
}

func (m *manager) Change(product *models.Product) (*models.Product, error) {
	return nil, nil
}

func (m *manager) Delete(id int) error {
	return nil
}

func (m *manager) Insert() error {
	return nil
}
