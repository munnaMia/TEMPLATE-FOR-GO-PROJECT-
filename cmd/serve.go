package cmd

import (
	"context"
	"log/slog"

	"github.com/munnaMia/ahlan/internal/config"
	"github.com/munnaMia/ahlan/internal/infra/auth"
	"github.com/munnaMia/ahlan/internal/infra/postgres"
	"github.com/munnaMia/ahlan/internal/usecase"
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
	pool, err := postgres.NewConnection(ctx, cnf)
	if err != nil {
		slog.Error("Error while initializing db connection", "err", err)
		return
	}
	defer pool.Close()

	// initializing services
	jwtService := auth.NewJWTService(cnf.Service.SecretKey)

	// initializing an http responder for http response.
	httpResponder := response.NewHttpResponse()

	// initializing the repositoris
	userRepo := postgres.NewUserRepository(pool)

	//initializing useCases or services
	userUsecase := usecase.NewUserUsecase(userRepo, jwtService)

	// initializing handlers
	userHandler := user.NewHandler(userUsecase, httpResponder)

	// creating a rest server.
	server := rest.NewServer(
		cnf,
		userHandler,
	)

	// start running the rest server.
	server.Start()
}
