package api

import "errors"

var (
	ErrValidation = errors.New("service validation error")
	ErrNotFound   = errors.New("service not found error")
)
