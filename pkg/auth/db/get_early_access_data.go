package auth

import (
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetEarlyAccessData(email string) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("early-access")

	var result authDomain.AddEarlyAccessParams

	err := collection.FindOne(*ctx, bson.D{{Key: "email", Value: email}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return nil
	}

	return err
}
