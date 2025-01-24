package users

import (
	"errors"
	"time"

	"github.com/qzich/orgserv/pkg/uuid"
)

type User struct {
	value *struct {
		ID        uuid.UUID
		Name      string
		Email     string
		Kind      KindEnum
		PassHash  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
}

func NewUser(id uuid.UUID, name string, email string, kind KindEnum, passHash string, createdAt time.Time, updatedAt time.Time) (newUser User, err error) {
	if id.IsZero() {
		return User{}, errors.New("id is required")
	}

	if err := Name(name).Validate(); err != nil {
		return User{}, err
	}

	if err := Email(email).Validate(); err != nil {
		return User{}, err
	}

	if err := kind.Validate(); err != nil {
		return User{}, err
	}

	if len(passHash) == 0 {
		return User{}, PassHashIsNotCorrect
	}

	if createdAt.IsZero() || updatedAt.IsZero() {
		return User{}, errors.New("created at and updated at are required")
	}

	newUser.value = &struct {
		ID        uuid.UUID
		Name      string
		Email     string
		Kind      KindEnum
		PassHash  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}{
		ID:        id,
		Name:      name,
		Email:     email,
		Kind:      kind,
		PassHash:  passHash,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	return
}

func (u User) IsZero() bool {
	return u.value == nil
}

func (u User) ID() uuid.UUID {
	return u.value.ID
}

func (u User) Name() string {
	return u.value.Name
}

func (u User) Email() string {
	return u.value.Email
}

func (u User) Kind() KindEnum {
	return u.value.Kind
}

func (u User) PasswordHash() string {
	return u.value.PassHash
}

func (u User) CreatedAt() time.Time {
	return u.value.CreatedAt
}

func (u User) UpdatedAt() time.Time {
	return u.value.UpdatedAt
}

// func (u *User) SetName(name string) {
// 	// u.user.Name = name
// 	user := *u.value
// 	user.Name = name
// 	u.value = &user
// }
