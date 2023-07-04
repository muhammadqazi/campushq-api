package middlewares

import (
	"context"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"

	"github.com/campushq-official/campushq-api/src/internal/common/logs"
	"github.com/campushq-official/campushq-api/src/internal/common/response"
	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	"github.com/campushq-official/campushq-api/src/internal/common/utils"
	"github.com/campushq-official/campushq-api/src/internal/config"
)

type Auth0Middleware struct {
	config config.Config
	log    *logs.Logger
}

func NewAuth0Middleware(config config.Config, log *logs.Logger) Auth0Middleware {
	return Auth0Middleware{
		config: config,
		log:    log,
	}
}

type contextKey string

const (
	TokenInfoContextKey contextKey = "decodedToken"
)

func (a *Auth0Middleware) Auth0Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		issuerURL, err := url.Parse("https://" + a.config.AUTH0_DOMAIN + "/")
		if err != nil {
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		}

		provider := jwks.NewCachingProvider(issuerURL, 5*time.Minute)

		jwtValidator, err := validator.New(
			provider.KeyFunc,
			validator.RS256,
			issuerURL.String(),
			[]string{a.config.AUTH0_DEVELOPEMENT_AUDIENCE},
		)
		if err != nil {
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		}

		authHeader := r.Header.Get("Authorization")
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 {
			response.JSONMessageResponse(rw, http.StatusUnauthorized, "Invalid authorization header")
			a.log.PrintHTTPResponse(r, http.StatusUnauthorized, "Invalid authorization header")
			return
		}

		_, err = jwtValidator.ValidateToken(r.Context(), authHeaderParts[1])
		if err != nil {
			response.JSONMessageResponse(rw, http.StatusUnauthorized, "Invalid token")
			a.log.PrintHTTPResponse(r, http.StatusUnauthorized, "Invalid token")
			return
		}

		/*
			"""
			Decode the token and add it to the request context
			"""
		*/

		decodedToken := utils.JWTExtractor(authHeaderParts[1])

		ctx := context.WithValue(r.Context(), TokenInfoContextKey, decodedToken)

		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
