package users

import (
	"context"

	"github.com/qzich/orgserv/entity/users"
)

type usersServiceClient struct{}

func (c usersServiceClient) CreateUser(ctx context.Context, name string, email string) (users.User, error) {
	panic("not implemented") // TODO: Implement
}
