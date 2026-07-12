package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

// log information about client request.
func (mdlw *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		slog.Info("http request",
			"method", r.Method,
			"path", r.URL.Path,
			"ip", r.RemoteAddr,
			"duration", time.Since(start),
		)
	})
}
