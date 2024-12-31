package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteTemporalCode(code int) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users-email-verification")

	collection.DeleteOne(*ctx, bson.D{{Key: "code", Value: code}})
}
