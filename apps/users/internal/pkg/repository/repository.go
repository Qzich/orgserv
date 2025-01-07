package repository

import (
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/uuid"
)

type UsersRepository interface {
	InsertUser(data users.User) error
	UpdateUser(userID uuid.UUID, data users.User) error
}
