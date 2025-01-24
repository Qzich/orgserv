package dto

import "time"

type AuthUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required" writeonly:"true"`
}

type UserDTO struct {
	ID       string `json:"id" readonly:"true"`
	Name     string `json:"name" validate:"required,min=4,max=255"`
	Email    string `json:"email" validate:"required,email"`
	Kind     string `json:"kind" validate:"required"`
	Password string `json:"password" validate:"required" writeonly:"true"`
	// PasswordHash string    `json:"password_hash" readonly:"true"`
	CreatedAt time.Time `json:"created_at" readonly:"true"`
	UpdatedAt time.Time `json:"updated_at" readonly:"true"`
}
