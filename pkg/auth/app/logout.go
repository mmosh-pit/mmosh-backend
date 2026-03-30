package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
)

func Logout(userId, token string) error {
	return auth.RemoveUserToken(userId, token)
}
