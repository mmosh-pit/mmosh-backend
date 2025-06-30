package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveEarlyAccess(params auth.AddEarlyAccessParams) error {

	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("early-access")

	_, err := collection.InsertOne(*ctx, params)

	return err
}
