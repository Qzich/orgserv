package service

import (
	"context"
	"fmt"
	"time"

	"github.com/qzich/orgserv/apps/users/internal"
	"github.com/qzich/orgserv/apps/users/internal/pkg/repository"
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg"
	"github.com/qzich/orgserv/pkg/api"
	"github.com/qzich/orgserv/pkg/uuid"
)

type usersService struct {
	repo repository.UsersRepository
}

func NewUserService(repo repository.UsersRepository) usersService {
	return usersService{repo: repo}
}

func (c usersService) AuthenticateUser(ctx context.Context, email string, password string) (users.User, error) {
	if err := users.Email(email).Validate(); err != nil {
		return users.User{}, err
	}

	// TOOD: add password specific validation rules and other error
	if len(password) == 0 {
		return users.User{}, fmt.Errorf("password is incorrect: %w", api.ErrValidation)
	}

	authUser, err := c.repo.GetAuthUser(email)
	if err != nil {
		return users.User{}, err
	}

	if !c.authenticate(authUser, password) {
		// TODO: add auth error
		return users.User{}, fmt.Errorf("authentication is failed: %w", api.ErrValidation)
	}

	return authUser.User(), nil
}

func (c usersService) CreateUser(ctx context.Context, name string, email string, kindStr string, password string) (users.User, error) {
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

	// TOOD: add password specific validation rules and other error
	if len(password) == 0 {
		return users.User{}, fmt.Errorf("password is incorrect: %w", api.ErrValidation)
	}

	passHash := c.hashPassword(password)

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

	return user, c.repo.InsertUser(
		pkg.Must(
			internal.NewAuthUser(user, passHash),
		),
	)
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

func (c usersService) hashPassword(password string) string {
	// TODO: do hash function from password here
	return "###"
}

func (c usersService) authenticate(authUser internal.AuthUser, pass string) bool {
	return authUser.PasswordHash() == c.hashPassword(pass)
}
