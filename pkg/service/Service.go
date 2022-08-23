package service

import (
	"github.com/mrbelka12000/artforintrovert_testEx/internal/repository"
	"github.com/mrbelka12000/artforintrovert_testEx/models"
)

type Service struct {
	models.Store
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Store: newProduct(repo),
	}
}
