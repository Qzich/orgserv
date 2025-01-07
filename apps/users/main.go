package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/qzich/orgserv/apps/users/internal/api/controller"
	"github.com/qzich/orgserv/apps/users/internal/api/router"
	"github.com/qzich/orgserv/apps/users/internal/pkg/repository/mysql"
	"github.com/qzich/orgserv/apps/users/internal/pkg/service"
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/api/json"
	logger "github.com/qzich/orgserv/pkg/logger/impl"
	pkgservice "github.com/qzich/orgserv/pkg/service"
	"github.com/qzich/orgserv/pkg/uuid"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log := logger.New()
	api := json.Api{}

	db, err := sql.Open("mysql", "root:roo@tcp(127.0.0.1:3306)/orgserv?parseTime=true")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	defer db.Close()

	userRepo := mysql.NewUsersRepository(db)
	userService := service.NewUserService(userRepo)
	usersCtl := controller.NewUser(log, api, api, userService)
	router := router.New(usersCtl.CreateUser, usersCtl.UsersList, usersCtl.GetUser)
	ctx := context.Background()

	log.Info(ctx, "Run users service")

	//createUser(ctx, userService)
	fmt.Println(getUsers(db))

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Info(ctx, "server closed\n")
		} else {
			log.Info(ctx, fmt.Sprintf("error starting server: %s\n", err))
			os.Exit(1)
		}
	}
}

func getUsers(db *sql.DB) []users.User {
	type userDAO struct {
		ID        int64
		UserId    string
		Name      string // required, min 4, max 255
		Email     string // required, email format
		Kind      string // required, support, customer
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	var res []users.User
	rows, err := db.Query("SELECT id, user_id, name, email, kind, created_at, updated_at FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user userDAO
		if err := rows.Scan(&user.ID, &user.UserId, &user.Name, &user.Email, &user.Kind, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}

		userID, err := uuid.FromString(user.UserId)
		if err != nil {
			panic(err)
		}

		res = append(res, users.User{
			ID:        userID,
			Name:      user.Name,
			Email:     user.Email,
			Kind:      user.Kind,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return res
}

func createUser(ctx context.Context, pkgservice pkgservice.UsersService) {
	_, err := pkgservice.CreateUser(ctx, "first", "qzichs@gmail.com", "support")
	if err != nil {
		panic(err)
	}
}
