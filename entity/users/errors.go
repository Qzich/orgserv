package users

import (
	"fmt"

	"github.com/qzich/orgserv/pkg/api"
)

var (
	ErrEmailIsNotCorrect = fmt.Errorf("email is incorrect: %w", api.ErrValidation)
	ErrKindIsNotCorrect  = fmt.Errorf("kind is incorrect: %w", api.ErrValidation)
	ErrNameIsNotCorrect  = fmt.Errorf("name length is incorrect: %w", api.ErrValidation)
)
