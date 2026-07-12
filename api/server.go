package api

import (
	"log/slog"
	"net/http"
	"os"
	"strconv"

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

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to ahlan"))
	})

	// prepare the resources.
	addr := ":" + strconv.Itoa(svr.cnf.Service.HTTP_Port)

	// http server.
	httpServer := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	slog.Info("Server start running", "PORT", svr.cnf.Service.HTTP_Port)
	if err := httpServer.ListenAndServe(); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
