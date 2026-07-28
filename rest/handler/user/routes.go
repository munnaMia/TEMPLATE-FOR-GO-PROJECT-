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
	mux.Handle("POST /api/auth/login", mngr.With(
		http.HandlerFunc(h.loginUser),
	))
	mux.Handle("POST /api/auth/logout", mngr.With(
		http.HandlerFunc(h.logoutUser),
	))
	mux.Handle("POST /api/auth/refresh-token", mngr.With(
		http.HandlerFunc(h.refreshToken),
	))

	// Personal Account & Profile Routes
	mux.Handle("GET /api/account/me", mngr.With(
		http.HandlerFunc(h.getCurrentUser),
	))
	mux.Handle("PATCH /api/account/me", mngr.With(
		http.HandlerFunc(h.updateCurrentUser),
	))
	mux.Handle("DELETE /api/account/me", mngr.With(
		http.HandlerFunc(h.deleteCurrentUser),
	))

	// Account Security & Recovery Routes
	mux.Handle("PUT /api/account/me/password", mngr.With(
		http.HandlerFunc(h.updatePassword),
	))
	mux.Handle("POST /api/account/forgot-password", mngr.With(
		http.HandlerFunc(h.forgotPassword),
	))
	mux.Handle("POST /api/account/reset-password", mngr.With(
		http.HandlerFunc(h.resetPassword),
	))
}
