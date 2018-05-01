package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger is middleware that logs request
func Logger(next http.Handler, name string) http.Handler {
	// returns handler function
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// start time
		start := time.Now()

		// run the handler
		next.ServeHTTP(w, r)

		// log the request details
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
