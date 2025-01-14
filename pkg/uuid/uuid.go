package uuid

import "github.com/gofrs/uuid"

type UUID struct {
	uuid *uuid.UUID `validate:"required"`
}

func New() UUID {
	uuid := uuid.Must(uuid.NewV4())
	return UUID{&uuid}
}

func FromString(s string) (UUID, error) {
	u, err := uuid.FromString(s)
	if err != nil {
		return UUID{}, err
	}

	return UUID{&u}, nil
}

func (u UUID) IsNil() bool {
	return u.uuid == nil
}

func (u UUID) String() string {
	return u.uuid.String()
}
