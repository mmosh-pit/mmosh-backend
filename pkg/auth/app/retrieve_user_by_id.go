package auth

import (
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
)

type IsAuthResponse struct {
	IsAuth bool             `json:"is_auth"`
	User   *authDomain.User `json:"user"`
}

func RetrieveUserById(id string) (IsAuthResponse, error) {
	user, err := authDb.GetUserById(id)

	if err != nil {
		return IsAuthResponse{IsAuth: false, User: nil}, err
	}

	return IsAuthResponse{IsAuth: true, User: &user}, nil
}
