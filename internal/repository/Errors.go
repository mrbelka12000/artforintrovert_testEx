package repository

import "errors"

var ErrNoData = errors.New("data not found")
var ErrContextTimeout = errors.New("bad connection")
var ErrInvalidId = errors.New("invalid id")
var ErrNoDocumentFound = errors.New("document not found")
var ErrUnknownError = errors.New("unknown error")
