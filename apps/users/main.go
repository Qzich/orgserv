package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/qzich/orgserv/apps/users/internal/api/controller"
	"github.com/qzich/orgserv/apps/users/internal/api/router"
	"github.com/qzich/orgserv/pkg/api/json"
	logger "github.com/qzich/orgserv/pkg/logger/impl"
)

func main() {
	log := logger.New()
	api := json.Api{}
	usersCtl := controller.NewUser(log, api, api)
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
