package entity

import (
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/api"
)

type AuthUser struct {
	value *struct {
		user     users.User
		passHash string
	}
}

func NewAuthUser(user users.User, passHash string) (authUser AuthUser, err error) {
	if user.IsZero() {
		return AuthUser{}, api.ErrValidation
	}

	if len(passHash) == 0 {
		return AuthUser{}, users.PassHashIsNotCorrect
	}

	authUser.value = &struct {
		user     users.User
		passHash string
	}{
		user:     user,
		passHash: passHash,
	}

	return
}

// func (a AuthUser) Authenticate(pass string) bool {
// 	// TODO: hash pass and compare with stored hash
// 	return true
// 	// return u.value.PassHash
// }

func (u AuthUser) IsZero() bool {
	return u.value == nil
}

func (a AuthUser) User() users.User {
	return a.value.user
}

func (a AuthUser) PasswordHash() string {
	return a.value.passHash
}
