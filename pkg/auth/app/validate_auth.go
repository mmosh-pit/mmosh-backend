package auth

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
)

func ValidateAuth(token string) (string, bool) {
	user, err := auth.GetUserBySessionToken(token)

	if err != nil {
		log.Printf("Error validating auth: %v\n", err)

		return "", false
	}

	return user.ID.Hex(), true
}
