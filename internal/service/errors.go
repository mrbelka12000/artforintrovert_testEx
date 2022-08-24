package service

import (
	"errors"
	"net/http"

	"go.uber.org/zap"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/service/repository"
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

func ParseErrorResponse(err error) (int, string) {
	zap.S().Debugf("error: %v", err)
	if errors.Is(err, ErrClientError) {
		return http.StatusBadRequest, err.Error()
	} else {
		return http.StatusInternalServerError, errors.Unwrap(err).Error()
	}
}
