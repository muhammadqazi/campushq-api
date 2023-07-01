package services

import (
	"context"

	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"

	"github.com/carlmjohnson/requests"
)

func (s *auth0Service) Auth0Signup(req dtos.StudentCreateDTO) error {

	ctx := context.TODO()

	token, _ := s.GenerateAccessToken()

	req.Connection = s.Connection

	err := requests.
		URL("/api/v2/users").
		Host(s.Domain).
		Header("Authorization", "Bearer "+token).
		BodyJSON(&req).
		Fetch(ctx)

	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
		return err
	}

	return nil

}
