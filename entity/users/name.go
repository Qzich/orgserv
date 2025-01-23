package users

import (
	"fmt"

	"github.com/qzich/orgserv/pkg/api"
)

var NameIsNotCorrect = fmt.Errorf("name length is incorrect: %w", api.ErrValidation)

type Name string

func (name Name) Validate() error {
	if len(name) < 4 || len(name) > 255 {
		return NameIsNotCorrect
	}

	return nil
}
