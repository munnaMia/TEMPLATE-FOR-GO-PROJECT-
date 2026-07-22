package response

import "net/http"

type Responder interface {
	Send(w http.ResponseWriter, status int, data any, meta *Meta)
	SendError(w http.ResponseWriter, status int, errCode, msg string, details []ValidationError)
}
