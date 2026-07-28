package cmd

import (
	"context"
	"log/slog"

	"github.com/munnaMia/ahlan/internal/config"
	"github.com/munnaMia/ahlan/internal/infra"
	"github.com/munnaMia/ahlan/internal/util/logger"
	"github.com/munnaMia/ahlan/internal/util/response"
	rest "github.com/munnaMia/ahlan/rest"
	"github.com/munnaMia/ahlan/rest/handler/user"
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
		slog.Error("Error while initializing db connection", "err", err)
		return
	}
	defer pool.Close()

	// initializing an http responder for http response.
	httpResponder := response.NewHttpResponse()

	// initializing handlers
	userHandler := user.NewHandler(httpResponder)

	// creating a rest server.
	restServer := rest.NewServer(
		cnf,
		userHandler,
	)

	// start running the rest server.
	restServer.Start()
}
