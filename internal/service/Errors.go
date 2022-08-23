package service

import (
	"errors"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/repository"
)

var (
	ErrClientError = errors.New("client error")
	ErrServerError = errors.New("server error")
)

func IsClientError(err error) bool {
	return errors.Is(err, repository.ErrInvalidId) ||
		errors.Is(err, repository.ErrNoDocumentFound)
}

func IsServerError(err error) bool {
	return errors.Is(err, repository.ErrNoData) ||
		errors.Is(err, repository.ErrContextTimeout)
}
