package middleware

import (
	"gorl/pkg/routers"
	"net/http"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authentication")

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		if r.URL.Path == "/admin" && len(token) == 0 {
			w.WriteHeader(http.StatusForbidden)
			routers.Handle403(wrapped, r)
			return
		}

		next.ServeHTTP(wrapped, r)
	})
}
