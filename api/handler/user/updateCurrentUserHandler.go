package user

import "net/http"

// Update profile fields (e.g., display name, avatar, bio).
func (h *Handler) updateCurrentUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's updateCurrentUser"))
}
