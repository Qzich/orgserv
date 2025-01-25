package entity

import (
	"fmt"

	"github.com/qzich/orgserv/pkg/api"
)

var (
	ErrPasswordIncorrect    = fmt.Errorf("password is incorrect: %w", api.ErrValidation)
	ErrAuthFailed           = fmt.Errorf("authentication is failed: %w", api.ErrValidation)
	ErrPassHashIsNotCorrect = fmt.Errorf("password hash length is incorrect: %w", api.ErrValidation)
)
