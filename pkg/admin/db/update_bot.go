package db

import (
	"log"

	adminDomain "github.com/mmosh-pit/mmosh_backend/pkg/admin/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateBot(payload adminDomain.UpdateBotPayload) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-project")

	_, err := collection.UpdateOne(*ctx, bson.D{{
		Key:   "_id",
		Value: payload.ID,
	}}, bson.M{
		"$set": bson.M{
			"name":         payload.Name,
			"symbol":       payload.Symbol,
			"defaultmodel": payload.DefaultModel,
			"deactivated":  payload.Deactivated,
		},
	})

	log.Printf("Error updating BOT: %v\n", err)
}
