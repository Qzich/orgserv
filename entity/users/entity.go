package users

import (
	"time"

	"github.com/qzich/orgserv/pkg/uuid"
)

type User struct {
	ID        uuid.UUID
	Name      string // required, min 4, max 255
	Email     string // required, email format
	Kind      Kind   // required, support, customer
	CreatedAt time.Time
	UpdatedAt time.Time
}
