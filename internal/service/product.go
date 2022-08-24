package service

import (
	"context"
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type productRepo struct {
	product ProductStoreRepo
}

func newProduct(product ProductStoreRepo) *productRepo {
	return &productRepo{
		product: product,
	}
}

func (p *productRepo) DeleteProduct(ctx context.Context, id string) error {
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

func (p *productRepo) GetAllProducts() ([]models.Product, error) {
	return p.product.GetAll()
}

func (p *productRepo) InsertProduct() error {
	return p.product.Insert()
}

func (p *productRepo) UpdateProduct(ctx context.Context, product *models.Product) error {
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
