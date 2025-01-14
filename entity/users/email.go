package users

import (
	"fmt"
	"net/mail"

	"github.com/qzich/orgserv/pkg/api"
)

type Email struct {
	*string `validate:"required"`
}

func NewEmail(email string) (Email, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return Email{}, fmt.Errorf("email is incorrect: %w", api.ErrValidation)
	}

	return Email{&email}, nil
}

func (e Email) Value() string {
	return *e.string
}
