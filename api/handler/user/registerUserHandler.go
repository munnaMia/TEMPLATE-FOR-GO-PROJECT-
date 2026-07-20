package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Register/Create a brand new user account.
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		h.responder.SendError(w, http.StatusBadRequest, "INVALID REQUEST BODY", "send a proper valid request data body", nil)
		return
	}

	fmt.Println(user)
	users = append(users, user)

	h.responder.Send(w, http.StatusOK, users, nil)
}
