package caserequests

import (
	"time"

	"github.com/gofrs/uuid"
	"github.com/qzich/orgserv/pkg/entity/assets"
	"github.com/qzich/orgserv/pkg/entity/users"
)

type Case struct {
	ID           uuid.UUID
	Number       string       // case number 8 digit with dashes
	Title        string       // min 4, max 255
	Description  string       // max 600
	Type         string       // ServiceRequest, InformationRequest
	SubType      string       // specific only for InformationRequest. Question, Notification
	Priority     string       // specific only for ServiceRequest. High, Normal, Low
	Origin       string       // options: API, CLI
	Asset        assets.Asset // Asset specific only for ServiceRequest
	Status       string       // New, Scheduled, In Progress, Verified, Done, Canceled
	StatusReason string       // max 255
	WorkReport   string       // specific only for ServiceRequest
	Comments     []string
	ClosedBy     users.User
	CreatedBy    users.User
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
