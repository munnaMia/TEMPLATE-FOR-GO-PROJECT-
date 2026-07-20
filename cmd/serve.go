package cmd

import (
	"log/slog"

	"github.com/munnaMia/ahlan/api"
	"github.com/munnaMia/ahlan/config"
	"github.com/munnaMia/ahlan/util/logger"
)

func Run() {
	// setting up new logger for app.
	newLogger := logger.NewLogger(false, false)
	slog.SetDefault(newLogger)

	// fetching the env's.
	cnf := config.GetConfiguration()

	// creating a api server.
	apiServer := api.NewServer(cnf)

	// start running the api server.
	apiServer.Start()
}
