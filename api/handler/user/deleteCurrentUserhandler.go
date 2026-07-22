package user

import "net/http"

// Deactivate or permanently soft-delete the logged-in user's account.
func (h *Handler) deleteCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's deleteCurrentUser"))
}
