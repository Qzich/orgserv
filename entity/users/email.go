package users

import (
	"fmt"
	"net/mail"

	"github.com/qzich/orgserv/pkg/api"
)

var EmailIsNotCorrect = fmt.Errorf("email is incorrect: %w", api.ErrValidation)

type Email string

func (e Email) Validate() error {
	if _, err := mail.ParseAddress(string(e)); err != nil {
		return EmailIsNotCorrect
	}

	return nil
}
