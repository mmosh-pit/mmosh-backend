package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveTemporalCode(email string, code int) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users-email-verification")

	collection.InsertOne(*ctx, map[string]interface{}{
		"email": email,
		"code":  code,
	})
}
