package services

import (
	"context"

	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	"github.com/carlmjohnson/requests"
)

var ApiTokenInfo struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
}

func (s *auth0Service) GenerateAccessToken() (string, error) {

	ctx := context.TODO()
	req := struct {
		GrantType    string `json:"grant_type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Audience     string `json:"audience"`
	}{
		GrantType:    "client_credentials",
		ClientID:     s.ClientID,
		ClientSecret: s.ClientSecret,
		Audience:     s.Audience,
	}

	err := requests.
		URL("/oauth/token").
		Host(s.Domain).
		BodyJSON(&req).
		ToJSON(&ApiTokenInfo).
		Fetch(ctx)

	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
		return "", err
	}

	return ApiTokenInfo.AccessToken, nil
}
