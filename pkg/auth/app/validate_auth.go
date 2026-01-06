package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
)

func ValidateAuth(token string, isAdmin bool) (string, bool) {
	user, err := auth.GetUserBySessionToken(token)

	if err != nil {
		return "", false
	}

	if user.Role != "wizard" && isAdmin {
		return "", false
	}

	return user.ID.Hex(), true
}
