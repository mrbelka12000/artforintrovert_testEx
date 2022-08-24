package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Product product
}

func NewRepository(client *mongo.Client) *Repository {
	return &Repository{
		Product: *newProduct(client),
	}
}
