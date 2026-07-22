package user

import "net/http"

// Authenticate credentials and return a token (JWT, cookie, etc.).
func (h *Handler) loginUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's loginUser"))
}
