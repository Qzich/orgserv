package service

import (
	"context"
	"fmt"
	"net/mail"
	"time"

	"github.com/qzich/orgserv/apps/users/internal/pkg/repository"
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/api"
	"github.com/qzich/orgserv/pkg/uuid"
)

type usersService struct {
	repo repository.UsersRepository
}

func NewUserService(repo repository.UsersRepository) usersService {
	return usersService{repo: repo}
}

func (c usersService) CreateUser(ctx context.Context, name string, email string, kindStr string) (users.User, error) {
	if len(name) < 4 || len(name) > 255 {
		return users.User{}, fmt.Errorf("name length is incorrect: %w", api.ErrValidation)
	}

	if _, err := mail.ParseAddress(email); err != nil {
		return users.User{}, fmt.Errorf("email is incorrect: %w", api.ErrValidation)
	}

	kind, err := users.KindEnumFromString(kindStr)
	if err != nil {
		return users.User{}, err
	}

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

func (c usersService) GetUser(ctx context.Context, userId uuid.UUID) (users.User, error) {
	return c.repo.GetUserByID(userId)
}

func (c usersService) AllUsers(ctx context.Context) ([]users.User, error) {
	searchUsers, err := c.repo.SearchUsers()
	if err != nil {
		return nil, err
	}

	if searchUsers == nil {
		searchUsers = make([]users.User, 0)
	}

	return searchUsers, nil
}
