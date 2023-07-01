package middlewares

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/auth0/go-jwt-middleware/v2/jwks"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/campushq-official/campushq-api/src/internal/common"
	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	"github.com/campushq-official/campushq-api/src/internal/config"
)

type Auth0Middleware struct {
	config config.Config
	log    *common.Logger
}

func NewAuth0Middleware(config config.Config, log *common.Logger) Auth0Middleware {
	return Auth0Middleware{
		config: config,
		log:    log,
	}
}

func (a *Auth0Middleware) Auth0TokenValidation(next http.Handler) http.Handler {
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
			[]string{a.config.AUTH0_AUDIENCE},
		)
		if err != nil {
			err = tracerr.Wrap(err)
			tracerr.PrintSourceColor(err)
		}

		// get the token from the request header
		authHeader := r.Header.Get("Authorization")
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 {
			a.log.PrintHTTPResponse(rw, r, http.StatusUnauthorized, "Invalid authorization header")
			return
		}

		tokenInfo, err := jwtValidator.ValidateToken(r.Context(), authHeaderParts[1])
		if err != nil {
			a.log.PrintHTTPResponse(rw, r, http.StatusUnauthorized, "Invalid token")
			return
		}

		fmt.Println(tokenInfo)

		// Go to next middleware:
		next.ServeHTTP(rw, r)
	})
}
