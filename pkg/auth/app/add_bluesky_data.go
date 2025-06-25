package auth

import (
	"log"

	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddBlueskyData(data authDomain.BlueskyUserData, userId string) error {
	userIdBson, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Printf("error transforming to Object ID bluesky: %v", err)
		return err
	}

	isValid := auth.IsValidConnection(data.Handle, data.Password)

	if !isValid {
		return authDomain.ErrInvalidBluesky
	}

	err = authDb.SaveBlueskyData(data, userIdBson)

	if err != nil {
		log.Printf("Got error saving bluesky data: %v\n", err)
		return err
	}

	return nil
}
