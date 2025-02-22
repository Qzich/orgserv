package repository

import (
	"github.com/qzich/orgserv/apps/users/internal/entity"
	"github.com/qzich/orgserv/apps/users/internal/pkg/password"
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/uuid"
)

type UsersRepository interface {
	InsertUser(data users.User, passHash password.Hash) error
	UpdateUser(userID uuid.UUID, data users.User) error
	GetUserByID(userID uuid.UUID) (users.User, error)
	GetAuthUser(email string) (entity.AuthUser, error)
	SearchUsers() ([]users.User, error)
}
