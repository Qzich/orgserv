package uuid

import "github.com/gofrs/uuid"

func New() UUID {
	return UUID{
		uuid.Must(uuid.NewV4()),
	}
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
