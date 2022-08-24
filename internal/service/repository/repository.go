// Package repository implements work with database.
package repository

import (
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/mrbelka12000/artforintrovert_testEx/pkg/logger"
)

type Repository struct {
	*product
	logger logger.Interface
}

func NewRepo(client *mongo.Client, l *logger.Logger) *Repository {
	return &Repository{
		product: newProduct(client, l),
	}
}
