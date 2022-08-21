package service

import (
	"github.com/mrbelka12000/artforintrovert_testEx/internal/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type Store interface {
	Change(product *models.Product) (*models.Product, error)
	GetAll() ([]models.Product, error)
	Delete(id int) error
	Insert() error
}

type Service struct {
	Store
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Store: newManager(repo),
	}
}
