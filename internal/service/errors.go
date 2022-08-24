package service

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/mrbelka12000/artforintrovert_testEx/internal/service/repository"
)

var (
	ErrClientError = errors.New("client error")
	ErrServerError = errors.New("server error")
)

func isClientError(err error) bool {
	return errors.Is(err, repository.ErrInvalidId) ||
		errors.Is(err, repository.ErrNoDocumentFound)
}

func isServerError(err error) bool {
	return errors.Is(err, repository.ErrNoData) ||
		errors.Is(err, repository.ErrContextTimeout)
}

func ParseErrorResponse(err error) (int, string) {
	if errors.Is(err, ErrClientError) {
		return http.StatusBadRequest, err.Error()
	} else {
		return http.StatusInternalServerError, errors.Unwrap(err).Error()
	}
}

func getErrorMsg(err error) error {
	if isClientError(err) {
		return fmt.Errorf("%w: %v", ErrClientError, err.Error())
	} else if isServerError(err) {
		return fmt.Errorf("%v: %w", err.Error(), ErrServerError)
	}
	return fmt.Errorf("%v: %w", err.Error(), errors.Unwrap(err))
}
