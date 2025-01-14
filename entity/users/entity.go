package users

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/qzich/orgserv/pkg/api"
	"github.com/qzich/orgserv/pkg/uuid"
)

type User struct {
	ID        uuid.UUID `validate:"required,uuid"`
	Name      string    `validate:"required"` // min 4, max 255
	Email     string    `validate:"required"` // email format
	Kind      KindEnum
	CreatedAt time.Time `validate:"required"`
	UpdatedAt time.Time `validate:"required"`
}

func NewUser(name string, email string, kind KindEnum) (User, error) {
	if len(name) < 4 || len(name) > 255 {
		return User{}, fmt.Errorf("name length is incorrect: %w", api.ErrValidation)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return User{}, fmt.Errorf("email is incorrect: %w", api.ErrValidation)
	}

	timeNow := time.Now().UTC()

	return User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Kind:      kind,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}, nil
}
