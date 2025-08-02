package middleware

import (
	"log"
	"net/http"
	"time"
)

type Logger interface {
	Info(next http.Handler) http.Handler
}

type Log struct{}

func (l Log) Info(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Request received: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

		// Call the next handler in the chain
		next.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("Request completed: %s %s took %s", r.Method, r.RequestURI, duration)
	})
}
