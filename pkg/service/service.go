package service

import (
	"context"
	"errors"

	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/uuid"
)

var (
	ErrUsersServiceError = errors.New("users service error")
)

type UsersService interface {
	AuthenticateUser(ctx context.Context, email string, password string) (users.User, error)
	CreateUser(ctx context.Context, name string, email string, kindStr string, password string) (users.User, error)
	GetUser(ctx context.Context, userId uuid.UUID) (users.User, error)
	AllUsers(ctx context.Context) ([]users.User, error)
}
