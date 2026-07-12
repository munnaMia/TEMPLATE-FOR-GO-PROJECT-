package http

import "github.com/munnaMia/ahlan/config"

type Server struct {
	cnf *config.Configuration
}

func NewServer(cnf *config.Configuration) *Server {
	return &Server{
		cnf: cnf,
	}
}
