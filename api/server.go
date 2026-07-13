package api

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"github.com/munnaMia/ahlan/api/handler/user"
	"github.com/munnaMia/ahlan/api/middleware"
	"github.com/munnaMia/ahlan/config"
)

type Server struct {
	cnf *config.Configuration
}

// return a new api server.
func NewServer(cnf *config.Configuration) *Server {
	return &Server{
		cnf: cnf,
	}
}

// start running the api server.
func (svr *Server) Start() {
	mux := http.NewServeMux()

	// initializing middleware and the middleware manager.
	mdlw := middleware.NewMiddleware()
	mdlwMngr := middleware.NewManager()

	// append global middlewares
	mdlwMngr.Use(
		mdlw.Logger,
	)

	// handlers
	userHandler := user.NewHandler()

	// register routes
	userHandler.RegisterRoutes(mux, mdlwMngr)

	// prepare the resources.
	addr := ":" + strconv.Itoa(svr.cnf.Service.HTTP_Port)
	wrapedMux := mdlwMngr.GlobalWraper(mux)

	// http server configuration.
	httpServer := &http.Server{
		Addr:    addr,
		Handler: wrapedMux,
	}

	slog.Info("Server start running", "PORT", svr.cnf.Service.HTTP_Port)
	if err := httpServer.ListenAndServe(); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
