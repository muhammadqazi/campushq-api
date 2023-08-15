package middlewares

import (
	"fmt"
	"net/http"
)

// CORS middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("CORS middleware")
		// Set CORS headers
		rw.Header().Set("Access-Control-Allow-Origin", "*")
		rw.Header().Set("Access-Control-Allow-Credentials", "true")
		rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == "OPTIONS" {
			rw.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println(rw.Header())

		// Call the next handler
		next.ServeHTTP(rw, r)
	})
}
