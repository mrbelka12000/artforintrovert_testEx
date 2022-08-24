package service

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type Product struct {
	product ProductStoreRepo
}

func NewProduct(product ProductStoreRepo) *Product {
	return &Product{
		product: product,
	}
}

func (p *Product) DeleteProduct(ctx context.Context, id string) error {
	err := p.product.Delete(ctx, id)
	if err != nil {
		if IsClientError(err) {
			return fmt.Errorf("%w: %v", ErrClientError, err.Error())
		} else if IsServerError(err) {
			return fmt.Errorf("%v: %w", err.Error(), ErrServerError)
		} else {
			zap.S().Debugf("unknown error received: %v", err)
			return fmt.Errorf("%v: %w", err.Error(), errors.Unwrap(err))
		}
	}
	return nil
}

func (p *Product) GetAllProducts() ([]models.Product, error) {
	return p.product.GetAll()
}

func (p *Product) InsertProduct() error {
	return p.product.Insert()
}

func (p *Product) UpdateProduct(ctx context.Context, product *models.Product) error {
	err := p.product.Update(ctx, product)
	if err != nil {
		if IsClientError(err) {
			return fmt.Errorf("%w: %v", ErrClientError, err.Error())
		} else if IsServerError(err) {
			return fmt.Errorf("%v: %w", err.Error(), ErrServerError)
		} else {
			zap.S().Debugf("unknown error received: %v", err)
			return fmt.Errorf("%v: %w", err.Error(), errors.Unwrap(err))
		}
	}
	return nil
}
