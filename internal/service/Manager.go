package service

import (
	"github.com/mrbelka12000/artforintrovert_testEx/internal/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type manager struct {
	repo *repository.Repository
}

func newManager(repo *repository.Repository) *manager {
	return &manager{repo}
}

func (m *manager) GetAll() ([]models.Product, error) {
	return m.repo.GetAll()
}

func (m *manager) Change(product *models.Product) (*models.Product, error) {
	return nil, nil
}

func (m *manager) Delete(id int) error {
	return nil
}

func (m *manager) Insert() error {
	return nil
}
