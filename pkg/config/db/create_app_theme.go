package db

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	configDomain "github.com/mmosh-pit/mmosh_backend/pkg/config/domain"
)

func CreateAppTheme(data configDomain.AppTheme) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("themes")

	collection.InsertOne(*ctx, data)
}
