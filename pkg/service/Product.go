package service

import (
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type product struct {
	repo *repository.Repository
}

func newProduct(repo *repository.Repository) *product {
	return &product{repo}
}

func (m *product) Delete(id string) error {
	err := m.repo.Delete(id)
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

func (m *product) GetAll() ([]models.Product, error) {
	return m.repo.GetAll()
}

func (m *product) Insert() error {
	return m.repo.Insert()
}

func (m *product) Update(product *models.Product) error {
	err := m.repo.Update(product)
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
