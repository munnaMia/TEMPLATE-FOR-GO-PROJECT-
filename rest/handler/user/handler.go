package user

import (
	"encoding/json"
	"net/http"

	"github.com/munnaMia/ahlan/internal/usecase"
	"github.com/munnaMia/ahlan/internal/util/response"
)

type Handler struct {
	usecase   *usecase.UserUsecase
	responder response.Responder
}

// return a new user handler.
func NewHandler(uc *usecase.UserUsecase, responder response.Responder) *Handler {
	return &Handler{
		usecase:   uc,
		responder: responder,
	}
}

// Register/Create a brand new user account.
func (h *Handler) registerUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.responder.SendError(w, http.StatusBadRequest, "INVALID REQUEST BODY", "Invalid JSON payload", nil)
		return
	}
}
