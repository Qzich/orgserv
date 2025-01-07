package service

import (
	"context"
	"time"

	"github.com/qzich/orgserv/apps/users/internal/pkg/repository"
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/uuid"
)

func NewUserService(repo repository.UsersRepository) usersService {
	return usersService{repo: repo}
}

type usersService struct {
	repo repository.UsersRepository
}

func (c usersService) CreateUser(ctx context.Context, name string, email string, kind string) (users.User, error) {
	timeNow := time.Now().UTC()
	user := users.User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Kind:      kind,
		UpdatedAt: timeNow,
		CreatedAt: timeNow,
	}

	return user, c.repo.InsertUser(user)
}
