package user

import "net/http"

// Change the current password (requires providing the old password).
func (h *Handler) updatePassword(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's updatePassword"))
}
