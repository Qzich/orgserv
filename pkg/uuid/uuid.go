package uuid

import "github.com/gofrs/uuid"

func New() UUID {
	return UUID{
		uuid.Must(uuid.NewV4()),
	}
}

func FromString(s string) (UUID, error) {
	u, err := uuid.FromString(s)
	if err != nil {
		return UUID{}, err
	}

	return UUID{u}, nil
}

type UUID struct {
	uuid uuid.UUID
}

func (u UUID) IsNil() bool {
	return u.uuid.IsNil()
}

func (u UUID) String() string {
	return u.uuid.String()
}
