package db

import (
	"log"

	adminDomain "github.com/mmosh-pit/mmosh_backend/pkg/admin/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(payload adminDomain.UpdateUserPayload) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	_, err := collection.UpdateOne(*ctx, bson.D{{
		Key:   "_id",
		Value: payload.ID,
	}}, bson.M{
		"$set": bson.M{
			"name":        payload.Name,
			"email":       payload.Email,
			"role":        payload.Role,
			"deactivated": payload.Deactivated,
		},
	})

	log.Printf("Error updating USER: %v\n", err)
}
