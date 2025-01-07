package service

import (
	"context"
	"errors"

	"github.com/qzich/orgserv/entity/users"
)

var (
	ErrUsersServiceError = errors.New("users service error")
)

type UsersService interface {
	CreateUser(ctx context.Context, name string, email string, kind string) (users.User, error)
}
