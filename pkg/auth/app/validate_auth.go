package auth

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
)

func ValidateAuth(token string, isAdmin bool) (string, bool) {
	user, err := auth.GetUserBySessionToken(token)

	if err != nil {
		log.Printf("Not authorized because of error: %v\n", err)

		return "", false
	}

	if user.Role != "wizard" && isAdmin {
		log.Println("Not authorized because not admin")
		return "", false
	}

	go auth.SaveLastLogin(user.ID)

	return user.ID.Hex(), true
}
