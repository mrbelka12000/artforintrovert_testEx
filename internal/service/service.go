package service

import (
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type ProductStore interface {
	Update(product *models.Product) error
	GetAll() ([]models.Product, error)
	Delete(id string) error
	Insert() error
}

type Service struct {
	Product *product
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Product: newProduct(&repo.Product),
	}
}
