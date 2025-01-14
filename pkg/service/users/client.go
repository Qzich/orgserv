package users

import (
	"context"

	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/api/dto"
)

type usersServiceClient struct{}

func (c usersServiceClient) CreateUser(ctx context.Context, nameStr string, emailStr string, kindStr string) (users.User, error) {
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

	var userDTO dto.UserDTO

	userDTO.Name = name.Value()
	userDTO.Email = email.Value()
	userDTO.Kind = kind.String()
	// TOOD: send json payload to create user endpoint
	_ = userDTO

	return users.NewUser(name, email, kind)
}
