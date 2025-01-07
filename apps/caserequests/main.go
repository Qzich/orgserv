package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/qzich/orgserv/apps/caserequests/internal/api/router"
	logger "github.com/qzich/orgserv/pkg/logger/impl"
)

func main() {
	log := logger.New()
	router := router.New()
	ctx := context.Background()

	log.Info(ctx, "Run caserequests service")

	if err := http.ListenAndServe(":8080", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			log.Info(ctx, "server closed\n")
		} else {
			log.Info(ctx, fmt.Sprintf("error starting server: %s\n", err))
			os.Exit(1)
		}
	}
}
