package service

import (
	"context"
	"time"

	"github.com/qzich/orgserv/apps/users/internal/pkg/repository"
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/uuid"
)

type usersService struct {
	repo repository.UsersRepository
}

func NewUserService(repo repository.UsersRepository) usersService {
	return usersService{repo: repo}
}

func (c usersService) CreateUser(ctx context.Context, name string, email string, kindStr string) (users.User, error) {
	if err := users.Name(name).Validate(); err != nil {
		return users.User{}, err
	}

	if err := users.Email(email).Validate(); err != nil {
		return users.User{}, err
	}

	kind, err := users.ParseKindFromString(kindStr)
	if err != nil {
		return users.User{}, err
	}

	timeNow := time.Now().UTC()

	user, err := users.NewUser(
		uuid.New(),
		name,
		email,
		kind,
		timeNow,
		timeNow,
	)
	if err != nil {
		return users.User{}, err
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
