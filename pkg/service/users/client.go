package users

import (
	"context"
	"time"

	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/api/dto"
	"github.com/qzich/orgserv/pkg/uuid"
)

type usersServiceClient struct{}

func (c usersServiceClient) CreateUser(ctx context.Context, name string, email string, kindStr string, passHash string) (users.User, error) {
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

	var userDTO dto.UserDTO

	userDTO.Name = name
	userDTO.Email = email
	userDTO.Kind = kind.String()
	// TOOD: send json payload to create user endpoint
	_ = userDTO

	timeNow := time.Now().UTC()
	userId := uuid.New()

	return users.NewUser(userId, name, email, kind, passHash, timeNow, timeNow)
}
