package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		log.Printf("INFO: \"%s %s %s\" %d \"%s\" %s",
			r.Method, r.URL.Path, r.Proto, wrapped.statusCode,
			r.UserAgent(), time.Since(start))
	})
}
