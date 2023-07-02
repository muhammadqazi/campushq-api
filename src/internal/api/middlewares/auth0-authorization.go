package middlewares

import (
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	"github.com/golang-jwt/jwt"
)

func Auth0Authorization(requiredPermission string, log *logs.Logger, handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		decodedToken := r.Context().Value(TokenInfoContextKey)
		if decodedToken == nil {
			log.PrintHTTPResponse(rw, r, http.StatusForbidden, "Invalid scope to access the resource.", false)
			rw.WriteHeader(http.StatusForbidden)
			return
		}

		claimsToken := decodedToken.(jwt.MapClaims)
		permissions := claimsToken["permissions"].([]interface{})

		for _, permission := range permissions {
			if permission == requiredPermission {
				handler(rw, r)
				return
			}
		}

		log.PrintHTTPResponse(rw, r, http.StatusForbidden, "Invalid scope to access the resource.", false)
		rw.WriteHeader(http.StatusForbidden)
	}
}
