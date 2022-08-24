package service

import (
	"context"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/models"
	"github.com/mrbelka12000/artforintrovert_testEx/pkg/logger"
)

type productRepo struct {
	product ProductStoreRepo
	l       logger.Interface
}

func newProduct(product ProductStoreRepo, log *logger.Logger) *productRepo {
	return &productRepo{
		product: product,
		l:       log,
	}
}

func (p *productRepo) DeleteProduct(ctx context.Context, id string) error {
	err := p.product.Delete(ctx, id)
	if err != nil {
		return getErrorMsg(err)
	}
	return nil
}

func (p *productRepo) GetAllProducts() ([]models.Product, error) {
	return p.product.GetAll()
}

func (p *productRepo) InsertProduct() error {
	return p.product.Insert()
}

func (p *productRepo) UpdateProduct(ctx context.Context, product *models.Product) error {
	err := p.product.Update(ctx, product)
	if err != nil {
		return getErrorMsg(err)
	}
	return nil
}
