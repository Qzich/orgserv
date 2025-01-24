package users

import "github.com/qzich/orgserv/pkg/api"

type AuthUser struct {
	value *struct {
		user     User
		passHash string
	}
}

func NewAuthUser(user User, passHash string) (authUser AuthUser, err error) {
	if user.IsZero() {
		return AuthUser{}, api.ErrValidation
	}

	if len(passHash) == 0 {
		return AuthUser{}, PassHashIsNotCorrect
	}

	authUser.value = &struct {
		user     User
		passHash string
	}{
		user:     user,
		passHash: passHash,
	}

	return
}

func (a AuthUser) Authenticate(pass string) bool {
	// TODO: hash pass and compare with stored hash
	return true
	// return u.value.PassHash
}

func (a AuthUser) User() User {
	return a.value.user
}
