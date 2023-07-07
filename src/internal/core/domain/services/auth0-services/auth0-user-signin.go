package services

import (
	"context"
	"fmt"

	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	"github.com/carlmjohnson/requests"
)

func (s *auth0Service) Auth0UserSignin(req dtos.StudentSigninDTO) (string, error) {
	ctx := context.TODO()

	req.Audience = s.DevelopementAudience
	req.ClientID = s.ClientID
	req.ClientSecret = s.ClientSecret
	req.GrantType = "password"
	req.Scope = "read:sample"
	username := fmt.Sprintf("student-%s", req.Username)
	req.Username = username

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
