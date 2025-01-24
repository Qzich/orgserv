package repository

import (
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/uuid"
)

type UsersRepository interface {
	InsertUser(data users.User, passHash string) error
	UpdateUser(userID uuid.UUID, data users.User) error
	GetUserByID(userID uuid.UUID) (users.User, error)
	GetAuthUser(email string) (users.AuthUser, error)
	SearchUsers() ([]users.User, error)
}
