package service

import (
	"context"

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

func (c usersService) CreateUser(ctx context.Context, nameStr string, emailStr string, kindStr string) (users.User, error) {
	kind, err := users.ParseKindFromString(kindStr)
	if err != nil {
		return users.User{}, err
	}

	email, err := users.NewEmail(emailStr)
	if err != nil {
		return users.User{}, err
	}

	name, err := users.NewName(nameStr)
	if err != nil {
		return users.User{}, err
	}

	user, err := users.NewUser(name, email, kind)
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
