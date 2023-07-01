package services

import (
	"context"
	"fmt"

	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	"github.com/carlmjohnson/requests"
)

func (s *auth0Service) GenerateAccessToken() (string, error) {

	ctx := context.TODO()
	req := struct {
		GrantType    string `json:"grant_type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Audience     string `json:"audience"`
	}{
		GrantType:    s.GrantType,
		ClientID:     s.ManagementClientID,
		ClientSecret: s.ManagementClientSecret,
		Audience:     s.Audience,
	}

	var res struct {
		AccessToken string `json:"access_token"`
	}
	err := requests.
		URL("/oauth/token").
		Host(s.Domain).
		BodyJSON(&req).
		ToJSON(&res).
		Fetch(ctx)

	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
		return "", err
	}

	fmt.Println(res)
	return res.AccessToken, nil

}
