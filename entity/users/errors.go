package users

import (
	"fmt"

	"github.com/qzich/orgserv/pkg/api"
)

var (
	EmailIsNotCorrect    = fmt.Errorf("email is incorrect: %w", api.ErrValidation)
	KindIsNotCorrect     = fmt.Errorf("kind is incorrect: %w", api.ErrValidation)
	NameIsNotCorrect     = fmt.Errorf("name length is incorrect: %w", api.ErrValidation)
	PassHashIsNotCorrect = fmt.Errorf("password hash length is incorrect: %w", api.ErrValidation)
)
