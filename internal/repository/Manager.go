package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/config"
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
		zap.S().Error("failed to get data")
		return nil, fmt.Errorf("failed to get data: %w", ErrNoData)
	}

	return tempData, nil
}

func (m *manager) Update(product *models.Product) error {
	cfg := config.GetConf()

	coll := m.client.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection)

	ctx, _ := context.WithTimeout(context.Background(), waitLimit)

	update := bson.D{{"$set", bson.D{{"name", product.Name}, {"price", product.Price}}}}
	result, err := coll.UpdateOne(ctx, bson.M{"_id": product.ID}, update)
	if err != nil {
		zap.S().Errorf("failed to delete: %v", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("failed to delete: %w", ErrContextTimeout)
		}
		return fmt.Errorf("%v: %w", err.Error(), ErrUnknownError)
	}

	if result.ModifiedCount != 1 {
		zap.S().Errorf("document with id %v does not exists", product.ID.String())
		return fmt.Errorf("%w", ErrNoDocumentFound)
	}

	needToUpdate = true
	return nil
}

func (m *manager) Delete(id string) error {
	cfg := config.GetConf()

	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		zap.S().Errorf("invalid id received %v ", id)
		return fmt.Errorf("%w", ErrInvalidId)
	}

	ctx, _ := context.WithTimeout(context.Background(), waitLimit)

	coll := m.client.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection)

	res, err := coll.DeleteOne(ctx, bson.M{"_id": primitiveId})
	if err != nil {
		zap.S().Errorf("failed to delete: %v", err)
		if errors.Is(err, context.DeadlineExceeded) {
			return fmt.Errorf("failed to delete: %w", ErrContextTimeout)
		}
		return fmt.Errorf("%v: %w", err.Error(), ErrUnknownError)
	}

	if res.DeletedCount != 1 {
		zap.S().Errorf("document with id %v does not exists", id)
		return fmt.Errorf("%w", ErrNoDocumentFound)
	}

	needToUpdate = true
	return nil
}

// Insert inserts default values to database.
func (m *manager) Insert() error {
	cfg := config.GetConf()

	coll := m.client.Database(cfg.MongoDB.Database).Collection(cfg.MongoDB.Collection)
	products := []models.Product{
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
