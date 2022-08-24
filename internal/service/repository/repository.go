package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repository struct {
	*product
}

func NewRepo(client *mongo.Client) *Repository {
	return &Repository{
		product: NewProduct(client),
	}
}
