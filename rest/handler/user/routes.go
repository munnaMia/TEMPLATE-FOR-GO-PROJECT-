package user

import (
	"net/http"

	"github.com/munnaMia/ahlan/rest/middleware"
)

// register user auth and accounts related routes and handlers
func (h *Handler) RegisterRoutes(mux *http.ServeMux, mngr *middleware.Manager) {

	// Authentication & Session Routes
	mux.Handle("POST /api/auth/register", mngr.With(
		http.HandlerFunc(h.registerUser),
	))
}
