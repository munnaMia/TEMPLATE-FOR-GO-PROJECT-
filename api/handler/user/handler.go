package user

import (
	"github.com/munnaMia/ahlan/util/response"
)

type Handler struct {
	responder response.Responder
}

// return a new user handler.
func NewHandler(responder response.Responder) *Handler {
	return &Handler{
		responder: responder,
	}
}

// when u start work on the route move then each of them on there separate file.

// only for now
type User struct {
	Name     string
	Email    string
	Password string
}

var users = make([]User, 0) // now database.
