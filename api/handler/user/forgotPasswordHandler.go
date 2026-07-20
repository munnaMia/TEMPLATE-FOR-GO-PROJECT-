package user

import "net/http"

// Request a password reset link/OTP sent to the user's email.
func (h *Handler) forgotPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's forgotPassword"))
}
