package services

import (
	"context"

	"github.com/campushq-official/campushq-api/src/internal/common/constants"
	"github.com/campushq-official/campushq-api/src/internal/common/tracerr"
	dtos "github.com/campushq-official/campushq-api/src/internal/core/domain/dtos/student-dtos"
	"github.com/carlmjohnson/requests"
)

func (s *auth0Service) Auth0AssignRole(userId string) (string, error) {
	ctx := context.TODO()

	var roles dtos.Auth0AssignRoleDTO

	studentRole := constants.Auth0Roles["STUDENT"]
	roles.Roles = []string{studentRole.ID}

	err := requests.
		URL("/api/v2/users/"+userId+"/roles").
		Host(s.Domain).
		Header("Authorization", "Bearer "+ApiTokenInfo.AccessToken).
		BodyJSON(&roles).
		Fetch(ctx)

	if err != nil {
		err = tracerr.Wrap(err)
		tracerr.PrintSourceColor(err)
		return "", err
	}

	return ApiTokenInfo.AccessToken, nil

}
