package middlewares

import (
	"net/http"
)

// CORS middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(rw, r)
	})
}
