package cmd

import (
	"context"
	"log/slog"

	"github.com/munnaMia/ahlan/api"
	"github.com/munnaMia/ahlan/api/handler/user"
	"github.com/munnaMia/ahlan/internal/config"
	"github.com/munnaMia/ahlan/internal/infra"
	"github.com/munnaMia/ahlan/internal/util/logger"
	"github.com/munnaMia/ahlan/internal/util/response"
)

func Run() {
	ctx := context.Background()

	// setting up new logger for app.
	newLogger := logger.NewLogger(false, false)
	slog.SetDefault(newLogger)

	// fetching the env's.
	cnf := config.GetConfiguration()

	// get a db connection
	pool, err := infra.NewConnection(ctx, cnf)
	if err != nil {
		// handle err carefully
	}
	defer pool.Close()

	// initializing an http responder for http response.
	httpResponder := response.NewHttpResponse()

	// initializing handlers
	userHandler := user.NewHandler(httpResponder)

	// creating a api server.
	apiServer := api.NewServer(
		cnf,
		userHandler,
	)

	// start running the api server.
	apiServer.Start()
}
