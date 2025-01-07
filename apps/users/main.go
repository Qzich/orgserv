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
	"github.com/qzich/orgserv/apps/users/internal/pkg/service"
	"github.com/qzich/orgserv/entity/users"
	"github.com/qzich/orgserv/pkg/api/json"
	logger "github.com/qzich/orgserv/pkg/logger/impl"

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

	//createUser(db)
	fmt.Println(getUsers(db))

	userService := service.NewUserService()
	usersCtl := controller.NewUser(log, api, api, userService)
	router := router.New(usersCtl.CreateUser, func(w http.ResponseWriter, r *http.Request) {}, func(w http.ResponseWriter, r *http.Request) {})
	ctx := context.Background()

	log.Info(ctx, "Run users service")

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
		Name      string // required, min 4, max 255
		Email     string // required, email format
		Kind      string // required, support, customer
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	var res []users.User
	rows, err := db.Query("SELECT id, name, email, kind, created_at, updated_at FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user userDAO
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Kind, &user.CreatedAt, &user.UpdatedAt); err != nil {
			panic(err)
		}
		res = append(res, users.User{
			Name:      user.Name,
			Email:     user.Email,
			Kind:      user.Kind,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return res
}

func createUser(db *sql.DB) {
	// var user users.User

	result, err := db.Exec("INSERT INTO users (name, email, kind) VALUES (?, ?, ?)", "first", "qzichs@gmail.com", "support")
	if err != nil {
		panic(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	fmt.Println(id)
	// user.ID = int(id)
	// user.CreatedAt = "now" // Placeholder
	// json.NewEncoder(w).Encode(user)
}
