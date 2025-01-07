package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/api"
	"github.com/qzich/orgserv/pkg/uuid"
)

func NewUsersRepository(db *sql.DB) usersRepository {
	return usersRepository{db: db}
}

type usersRepository struct {
	db *sql.DB
}

type userDAO struct {
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

func (r usersRepository) GetUserByID(userID uuid.UUID) (users.User, error) {
	var user userDAO
	err := r.db.QueryRow(
		"SELECT id, user_id, name, email, kind, created_at, updated_at FROM users WHERE user_id = ? LIMIT 1", userID.String(),
	).Scan(&user.ID, &user.UserID, &user.Name, &user.Email, &user.Kind, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return users.User{}, fmt.Errorf("repo has no rows: %w", api.ErrNotFound)
		}
		return users.User{}, err
	}

	userId, err := uuid.FromString(user.UserID)
	if err != nil {
		return users.User{}, err
	}

	return users.User{
		ID:        userId,
		Name:      user.Name,
		Email:     user.Email,
		Kind:      user.Kind,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil

	// var res []users.User
	// rows, err := db.Query("SELECT id, user_id, name, email, kind, created_at, updated_at FROM users")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var user userDAO
	// 	if err := rows.Scan(&user.ID, &user.UserId, &user.Name, &user.Email, &user.Kind, &user.CreatedAt, &user.UpdatedAt); err != nil {
	// 		panic(err)
	// 	}

	// 	userID, err := uuid.FromString(user.UserId)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	res = append(res, users.User{
	// 		ID:        userID,
	// 		Name:      user.Name,
	// 		Email:     user.Email,
	// 		Kind:      user.Kind,
	// 		CreatedAt: user.CreatedAt,
	// 		UpdatedAt: user.UpdatedAt,
	// 	})
	// }
	// return res

}
