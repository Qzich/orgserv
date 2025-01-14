package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg"
	"github.com/qzich/orgserv/pkg/api"
	"github.com/qzich/orgserv/pkg/storage"
	"github.com/qzich/orgserv/pkg/uuid"
	"github.com/qzich/orgserv/pkg/validate"
)

type (
	usersRepository struct {
		db *sql.DB
	}
	userDAO struct {
		ID        int64
		UserID    string
		Name      string
		Email     string
		Kind      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

func NewUsersRepository(connectionString string) (usersRepository, *sql.DB) {
	db, err := storage.NewMysqlConnection(connectionString)
	if err != nil {
		panic(err)
	}

	return usersRepository{db: db}, db
}

func (r usersRepository) InsertUser(data users.User) error {
	// validate whole struct in case if users.User was not initialized via constructor func
	if err := validate.Struct(data); err != nil {
		return err
	}

	_, err := r.db.Exec(
		"INSERT INTO users (user_id, name, email, kind, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		data.ID.String(),
		data.Name,
		data.Email,
		data.Kind.String(), // TODO: use kind value or id
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
	var dao userDAO
	err := r.db.QueryRow(
		"SELECT id, user_id, name, email, kind, created_at, updated_at FROM users WHERE user_id = ? LIMIT 1", userID.String(),
	).Scan(&dao.ID, &dao.UserID, &dao.Name, &dao.Email, &dao.Kind, &dao.CreatedAt, &dao.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return users.User{}, fmt.Errorf("repo has no rows: %w", api.ErrNotFound)
		}
		return users.User{}, err
	}

	userId, err := uuid.FromString(dao.UserID)
	if err != nil {
		return users.User{}, err
	}

	return users.User{
		ID:        userId,
		Name:      dao.Name,
		Email:     dao.Email,
		Kind:      pkg.Must(users.ParseKindFromString(dao.Kind)),
		CreatedAt: dao.CreatedAt,
		UpdatedAt: dao.UpdatedAt,
	}, nil
}

func (r usersRepository) SearchUsers() ([]users.User, error) {
	var res []users.User

	rows, err := r.db.Query("SELECT id, user_id, name, email, kind, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var dao userDAO
		if err := rows.Scan(&dao.ID, &dao.UserID, &dao.Name, &dao.Email, &dao.Kind, &dao.CreatedAt, &dao.UpdatedAt); err != nil {
			return nil, err
		}

		userID, err := uuid.FromString(dao.UserID)
		if err != nil {
			return nil, err
		}

		res = append(res, users.User{
			ID:        userID,
			Name:      dao.Name,
			Email:     dao.Email,
			Kind:      pkg.Must(users.ParseKindFromString(dao.Kind)),
			CreatedAt: dao.CreatedAt,
			UpdatedAt: dao.UpdatedAt,
		})
	}
	return res, nil
}
