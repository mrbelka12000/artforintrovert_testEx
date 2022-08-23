package service

import (
	"errors"
	"fmt"

	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type manager struct {
	repo *repository.Repository
}

func newManager(repo *repository.Repository) *manager {
	return &manager{repo}
}

func (m *manager) Delete(id string) error {
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

func (m *manager) GetAll() ([]models.Product, error) {
	return m.repo.GetAll()
}

func (m *manager) Insert() error {
	return m.repo.Insert()
}

func (m *manager) Update(product *models.Product) error {
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
