package user

import "net/http"

// Register/Create a brand new user account.
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's registerUser"))
}
