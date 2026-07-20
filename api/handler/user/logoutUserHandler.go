package user

import "net/http"

// Invalidate the current session token or clear session cookies.
func (h *Handler) logoutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's logoutUser"))
}
