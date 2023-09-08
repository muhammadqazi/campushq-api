package middlewares

import (
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/logs"
)

func Loggin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger := logs.NewLogger(req)
		logger.PrintHTTPRequest()
		next.ServeHTTP(w, req)
	})
}
