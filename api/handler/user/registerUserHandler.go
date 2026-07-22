package user

import (
	"encoding/json"
	"net/http"
)

// Register/Create a brand new user account.
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var user User
	users := make([]User, 0) // temp

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.responder.SendError(w, http.StatusBadRequest, "INVALID REQUEST BODY", "send a proper valid request data body", nil)
		return
	}

	h.responder.Send(w, http.StatusOK, users, nil)
}
