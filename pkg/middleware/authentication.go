package middleware

import (
	"net/http"
)

var adminOnly = []string{"/admin", "/api/v1/create_user"}
var authorizedOnly = []string{"/api/v1/create_random_link", "/api/v1/create_link"}

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// token := r.Header.Get("TOKEN")

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		// isAdminPath := slices.Contains(adminOnly, r.URL.Path)
		// isAuthorizedPath := slices.Contains(authorizedOnly, r.URL.Path)

		next.ServeHTTP(wrapped, r)
	})
}
