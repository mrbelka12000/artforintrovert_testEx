package repository

import (
	"github.com/mrbelka12000/artforintrovert_testEx/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Store interface {
	Update(product *models.Product) error
	GetAll() ([]models.Product, error)
	Delete(id string) error
	Insert() error
}

type Repository struct {
	Store
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		Store: newManager(client),
	}
}
