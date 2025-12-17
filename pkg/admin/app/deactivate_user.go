package app

import (
	adminDb "github.com/mmosh-pit/mmosh_backend/pkg/admin/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeactivateUser(userId string) error {
	parsedUserId, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return err
	}

	return adminDb.DeactivateUser(&parsedUserId)
}
