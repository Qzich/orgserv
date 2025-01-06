package uuid

import "github.com/gofrs/uuid"

type UUID struct {
	uuid uuid.UUID
}

func (u UUID) IsNil() bool {
	return u.uuid.IsNil()
}
