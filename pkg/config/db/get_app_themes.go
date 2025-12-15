package db

import (
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	configDomain "github.com/mmosh-pit/mmosh_backend/pkg/config/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAppThemes() *[]configDomain.AppTheme {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("themes")

	var result []configDomain.AppTheme

	res, err := collection.Find(*ctx, bson.D{{}}, nil)

	if err != nil {
		log.Printf("[APP THEMES] could not get themes: %v\n", err)
		return &result
	}

	for res.Next(*ctx) {
		var theme configDomain.AppTheme

		if err := res.Decode(&theme); err != nil {
			log.Printf("[APP THEMES] Got error decoding theme: %v\n", err)
			continue
		}

		result = append(result, theme)
	}

	return &result
}
