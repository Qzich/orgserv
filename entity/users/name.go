package users

import (
	"fmt"

	"github.com/qzich/orgserv/pkg/api"
)

type Name struct {
	*string `validate:"required"`
}

func NewName(name string) (Name, error) {
	if len(name) < 4 || len(name) > 255 {
		return Name{}, fmt.Errorf("name length is incorrect: %w", api.ErrValidation)
	}

	return Name{&name}, nil
}

func (n Name) Value() string {
	return *n.string
}
