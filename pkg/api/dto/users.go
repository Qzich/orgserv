package dto

import "time"

type UserDTO struct {
	ID        string    `json:"id" readonly:"true"`
	Name      string    `json:"name" validate:"required,min=4,max=255"`
	Email     string    `json:"email" validate:"required,email"`
	Kind      string    `json:"kind" validate:"required"`
	PassHash  string    `json:"pass_hash" validate:"required"`
	CreatedAt time.Time `json:"created_at" readonly:"true"`
	UpdatedAt time.Time `json:"updated_at" readonly:"true"`
}
