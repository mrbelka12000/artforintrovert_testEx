package repository

import "errors"

var (
	ErrNoData          = errors.New("data not found")
	ErrContextTimeout  = errors.New("bad connection")
	ErrInvalidId       = errors.New("invalid id")
	ErrNoDocumentFound = errors.New("document not found")
	ErrUnknownError    = errors.New("unknown error")
)
