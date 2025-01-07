package service

import (
	"context"
	"time"

	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/uuid"
)

func NewUserService() usersService {
	return usersService{}
}

type usersService struct{}

func (c usersService) CreateUser(ctx context.Context, name string, email string, kind string) (users.User, error) {
	timeNow := time.Now().UTC()
	return users.User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Kind:      kind,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}, nil
}
