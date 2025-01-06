package users

import (
	"time"

	"github.com/gofrs/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string // required, min 4, max 255
	Email     string // required, email format
	Kind      string // required, support, customer
	CreatedAt time.Time
	UpdatedAt time.Time
}
