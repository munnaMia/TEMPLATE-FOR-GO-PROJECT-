package user

import "net/http"

// Renew an expired access token using a refresh token.
func (h *Handler) refreshToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Now let's refreshToken"))
}
