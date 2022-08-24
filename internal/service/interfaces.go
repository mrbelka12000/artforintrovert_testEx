package service

import (
	"context"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type (
	ProductStoreSrv interface {
		UpdateProduct(ctx context.Context, product *models.Product) error
		GetAllProducts() ([]models.Product, error)
		DeleteProduct(ctx context.Context, id string) error
		InsertProduct() error
	}

	ProductStoreRepo interface {
		Update(ctx context.Context, product *models.Product) error
		GetAll() ([]models.Product, error)
		Delete(ctx context.Context, id string) error
		Insert() error
	}
)
