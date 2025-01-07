package mysql

import (
	"database/sql"
	"time"

	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/uuid"
)

func NewUsersRepository(db *sql.DB) usersRepository {
	return usersRepository{db: db}
}

type usersRepository struct {
	db *sql.DB
}

type user struct {
	ID        int64
	UserID    string
	Name      string
	Email     string
	Kind      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r usersRepository) InsertUser(data users.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (user_id, name, email, kind, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)",
		data.ID.String(),
		data.Name,
		data.Email,
		data.Kind,
		data.CreatedAt,
		data.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r usersRepository) UpdateUser(userID uuid.UUID, data users.User) error {
	panic("not implemented") // TODO: Implement
}
