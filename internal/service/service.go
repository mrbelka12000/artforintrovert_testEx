package service

import (
	"github.com/mrbelka12000/artforintrovert_testEx/internal/service/repository"
)

type Service struct {
	Product *product
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Product: newProduct(&repo.Product),
	}
}
