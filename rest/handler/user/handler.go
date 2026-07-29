package user

import (
	"net/http"

	"github.com/munnaMia/ahlan/internal/util/response"
)

type Handler struct {
	responder response.Responder
}

// return a new user handler.
func NewHandler(responder response.Responder) *Handler {
	return &Handler{
		responder: responder,
	}
}

// Renew an expired access token using a refresh token.
func (h *Handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's refreshToken"))
}

// Register/Create a brand new user account.
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {

}

// Authenticate credentials and return a token (JWT, cookie, etc.).
func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's loginUser"))
}

// Invalidate the current session token or clear session cookies.
func (h *Handler) logoutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's logoutUser"))
}

// Fetch the current logged-in user's profile info.
func (h *Handler) getCurrentUser(w http.ResponseWriter, r *http.Request) {

}

// Update profile fields (e.g., display name, avatar, bio).
func (h *Handler) updateCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's updateCurrentUser"))
}

// Deactivate or permanently soft-delete the logged-in user's account.
func (h *Handler) deleteCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's deleteCurrentUser"))
}

// Submit a new password alongside the token/OTP received.
func (h *Handler) resetPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's resetPassword"))
}

// Request a password reset link/OTP sent to the user's email.
func (h *Handler) forgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's forgotPassword"))
}

// Change the current password (requires providing the old password).
func (h *Handler) updatePassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's updatePassword"))
}
