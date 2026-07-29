package response

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

type HTTPResponse struct{}

// return a new http response struct instance.
func NewHttpResponse() Responder {
	return &HTTPResponse{}
}

type Meta struct {
	TotalCount int64 `json:"totalCount"`
	Limit      int   `json:"limit"`
	Offset     int   `json:"offset"`
}

type ValidationError struct {
	Field string `json:"field"`
	Issue string `json:"issue"`
}

type jsonEnvelope struct {
	Data any   `json:"data"`
	Meta *Meta `json:"meta"`
}

type errorEnvelope struct {
	Error struct {
		Code    string            `json:"code"`
		Message string            `json:"message"`
		Details []ValidationError `json:"details,omitempty"`
	} `json:"error"`
}

// Write json response to client
func (httpRes *HTTPResponse) Send(w http.ResponseWriter, statusCode int, data any, meta *Meta) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(jsonEnvelope{
		Data: data,
		Meta: meta,
	})

	if err != nil {
		slog.Error("Could not sending JSON response", "Err", err)
		return
	}
}

// Write json error response to client
func (httpRes *HTTPResponse) SendError(w http.ResponseWriter, statusCode int, errCode, msg string, details []ValidationError) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	res := errorEnvelope{}

	res.Error.Code = errCode
	res.Error.Message = msg
	res.Error.Details = details

	err := json.NewEncoder(w).Encode(res)

	if err != nil {
		slog.Error("Could not sending JSON response", "Err", err)
		return
	}
}
