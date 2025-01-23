package uuid

import (
	"github.com/gofrs/uuid"
	"github.com/qzich/orgserv/pkg"
)

type UUID struct{ value *uuid.UUID }

func New() UUID {
	value := pkg.Must(uuid.NewV4())
	return UUID{&value}
}

func FromString(s string) (UUID, error) {
	u, err := uuid.FromString(s)
	if err != nil {
		return UUID{}, err
	}

	return UUID{&u}, nil
}

func (u UUID) IsZero() bool {
	return u.value == nil
}

func (u UUID) String() string {
	return u.value.String()
}
