package users

import (
	"net/mail"
)

type Email string

func (e Email) Validate() error {
	if _, err := mail.ParseAddress(string(e)); err != nil {
		return ErrEmailIsNotCorrect
	}

	return nil
}
