package entity

import (
	"fmt"

	"github.com/qzich/orgserv/apps/users/internal/pkg/password"
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/api"
)

type AuthUser struct {
	value *struct {
		user     users.User
		passHash password.Hash
	}
}

func NewAuthUser(user users.User, passHash password.Hash) (authUser AuthUser, err error) {
	if user.IsZero() {
		return AuthUser{}, api.ErrValidation
	}

	if len(passHash) == 0 {
		return AuthUser{}, users.PassHashIsNotCorrect
	}

	authUser.value = &struct {
		user     users.User
		passHash password.Hash
	}{
		user:     user,
		passHash: passHash,
	}

	return
}

func (a AuthUser) Authenticate(verify func(password.Hash) bool) (users.User, error) {
	if verify(a.value.passHash) {
		return a.value.user, nil
	}

	return users.User{}, fmt.Errorf("authentication is failed: %w", api.ErrValidation)
}
