package assets

import (
	"time"

	"github.com/qzich/orgserv/pkg/uuid"
)

type Asset struct {
	ID          uuid.UUID
	Number      string // required, asset number ISO_CODE+6 digit
	Name        string // required, min 2, max 24
	Description string // optional, max 600
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
