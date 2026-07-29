package middleware

import "github.com/munnaMia/ahlan/internal/config"

type Middleware struct {
	cnf *config.Configuration
}

// return a new middleware struct pointer.
func NewMiddleware(cnf *config.Configuration) *Middleware {
	return &Middleware{
		cnf: cnf,
	}
}
