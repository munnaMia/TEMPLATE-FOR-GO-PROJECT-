package user

import "net/http"

// Submit a new password alongside the token/OTP received.
func (h *Handler) resetPassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's resetPassword"))
}
