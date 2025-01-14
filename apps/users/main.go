package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/qzich/orgserv/apps/users/internal/api/controller"
	"github.com/qzich/orgserv/apps/users/internal/api/router"
	"github.com/qzich/orgserv/apps/users/internal/pkg/repository/mysql"
	"github.com/qzich/orgserv/apps/users/internal/pkg/service"
	"github.com/qzich/orgserv/pkg/api/json"
	logger "github.com/qzich/orgserv/pkg/logger/impl"
)

func main() {
	log := logger.New()
	api := json.Api{}

	userRepo, db := mysql.NewUsersRepository("root:roo@tcp(127.0.0.1:3306)/orgserv?parseTime=true")
	defer db.Close()

	userService := service.NewUserService(userRepo)
	usersCtl := controller.NewUser(log, api, api, userService)
	router := router.New(usersCtl.CreateUser, usersCtl.UsersList, usersCtl.GetUser)
	ctx := context.Background()

	log.Info(ctx, "Run users service")

	// createUser(ctx, userService)
	// fmt.Println(getUsers(ctx, userService))

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Info(ctx, "server closed\n")
		} else {
			log.Info(ctx, fmt.Sprintf("error starting server: %s\n", err))
			os.Exit(1)
		}
	}
}

// func getUsers(ctx context.Context, pkgservice pkgservice.UsersService) []users.User {
// 	users, err := pkgservice.AllUsers(ctx)
// 	if err != nil {
// 		panic(err)
// 	}

// 	return users
// }

// func createUser(ctx context.Context, pkgservice pkgservice.UsersService) {
// 	_, err := pkgservice.CreateUser(ctx, "first", "qzichs@gmail.com", "support")
// 	if err != nil {
// 		panic(err)
// 	}
// }
