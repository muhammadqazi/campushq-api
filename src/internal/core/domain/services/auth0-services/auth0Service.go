package services

import (
	"github.com/campushq-official/campushq-api/src/internal/config"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
)

type Auth0Service interface {
	GenerateAccessToken() (string, error)
	Auth0Signup(req *dtos.Auth0SignupDTO) error
	Auth0UserSignin(req dtos.StudentSigninDTO) (string, error)
}

type auth0Service struct {
	Audience             string `json:"audience"`
	Domain               string `json:"domain"`
	Connection           string `json:"connection"`
	ClientID             string `json:"client_id"`
	ClientSecret         string `json:"client_secret"`
	DevelopementAudience string `json:"developement_audience"`
}

func NewAuth0Service(config config.Config) Auth0Service {
	return &auth0Service{
		ClientID:             config.AUTH0_CLIENT_ID,
		ClientSecret:         config.AUTH0_CLIENT_SECRET,
		Audience:             config.AUTH0_AUDIENCE,
		Domain:               config.AUTH0_DOMAIN,
		Connection:           config.AUTH0_CONNECTION,
		DevelopementAudience: config.AUTH0_DEVELOPEMENT_AUDIENCE,
	}
}
