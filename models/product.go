package models

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	minimumPrice = 0
)

type Product struct {
	ID    primitive.ObjectID `json:"_id" bson:"_id"`
	Name  string             `json:"name" bson:"name"`
	Price float64            `json:"price" bson:"price"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("can not update to empty name")
	}
	if p.Price <= minimumPrice {
		return errors.New("the price cannot be less than the minimum")
	}
	return nil
}
