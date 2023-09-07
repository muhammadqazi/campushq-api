package middlewares

import (
	"net/http"

	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	"github.com/campushq-official/campushq-api/src/internal/common/response"
	"github.com/golang-jwt/jwt"
)

func Auth0Authorization(requiredPermission string, log *logs.Logger, handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		decodedToken := r.Context().Value(TokenInfoContextKey)
		if decodedToken == nil {
			response.JSONMessageResponse(rw, http.StatusForbidden, "Invalid scope to access the resource.")
			log.PrintHTTPResponse(r, http.StatusForbidden, "Invalid scope to access the resource.")
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

		response.JSONMessageResponse(rw, http.StatusForbidden, "Invalid scope to access the resource.")
		log.PrintHTTPResponse(r, http.StatusForbidden, "Invalid scope to access the resource.")
	}
}
