package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type Repository struct {
	models.Store
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		Store: newProduct(client),
	}
}
