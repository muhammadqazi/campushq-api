package services

import "github.com/campushq-official/campushq-api/src/internal/config"

type Auth0Service interface {
	GenerateAccessToken() (string, error)
}

type auth0Service struct {
	GrantType              string `json:"grant_type"`
	Audience               string `json:"audience"`
	Domain                 string `json:"domain"`
	ManagementClientID     string `json:"client_id"`
	ManagementClientSecret string `json:"client_secret"`
}

func NewAuth0Service(config config.Config) Auth0Service {
	return &auth0Service{
		GrantType:              "client_credentials",
		ManagementClientID:     config.AUTH0_MANAGEMENT_CLIENT_ID,
		ManagementClientSecret: config.AUTH0_MANAGEMENT_CLIENT_SECRET,
		Audience:               config.AUTH0_AUDIENCE,
		Domain:                 config.AUTH0_DOMAIN,
	}
}
