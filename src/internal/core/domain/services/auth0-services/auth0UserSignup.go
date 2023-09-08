package services

import (
	"context"

	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"

	"github.com/carlmjohnson/requests"
)

func (s *auth0Service) Auth0Signup(req *dtos.Auth0SignupDTO) (*dtos.Auth0SignupResponseDTO, error) {

	var response *dtos.Auth0SignupResponseDTO
	ctx := context.TODO()

	token, _ := s.GenerateAccessToken()

	req.Connection = s.Connection

	err := requests.
		URL("/api/v2/users").
		Host(s.Domain).
		Header("Authorization", "Bearer "+token).
		BodyJSON(&req).
		ToJSON(&response).
		Fetch(ctx)

	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
		return nil, err
	}

	return response, nil

}
