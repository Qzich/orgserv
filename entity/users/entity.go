package users

import (
	"time"

	"github.com/qzich/orgserv/pkg/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      Name
	Email     Email
	Kind      KindEnum
	CreatedAt time.Time `validate:"required"`
	UpdatedAt time.Time `validate:"required"`
}

func NewUser(name Name, email Email, kind KindEnum) (User, error) {
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
