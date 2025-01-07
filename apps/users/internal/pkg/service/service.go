package service

import (
	"context"

	"github.com/qzich/orgserv/entity/users"
)

type usersService struct{}

func (c usersService) CreateUser(ctx context.Context, name string, email string) (users.User, error) {
	panic("not implemented") // TODO: Implement
}
