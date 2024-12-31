package auth

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Logout(userId, token string) error {

	objectId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Printf("Invalid Object ID %s\n", userId)
		return err
	}

	err = auth.RemoveUserToken(&objectId, token)

	return err
}
