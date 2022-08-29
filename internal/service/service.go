// Package service implements application business logic.
package service

import "github.com/mrbelka12000/artforintrovert_testEx/pkg/logger"

type Service struct {
	ProductStoreSrv
	l *logger.Logger
}

func NewService(repo ProductStoreRepo, log *logger.Logger) *Service {
	return &Service{
		ProductStoreSrv: newProduct(repo, log),
	}
}
