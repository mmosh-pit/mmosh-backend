package auth

import (
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetTemporalCode(code int) *authDomain.VerificationData {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users-email-verification")

	var result authDomain.VerificationData

	err := collection.FindOne(*ctx, bson.D{{Key: "code", Value: code}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return nil
	}

	return &result
}
