package notificationDb

import (
	"context"
	"fmt"
	"time"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InsertPlayerId(playerId string, wallet string) (bool, error) {
	client, _ := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-player")

	// Prepare update operations
	filter := bson.M{"wallet": wallet}
	update := bson.M{
		"$addToSet": bson.M{"playerIds": playerId}, // Add playerId only if not present
		"$set":      bson.M{"updated_date": time.Now().UTC()},
		"$setOnInsert": bson.M{
			"created_date": time.Now().UTC(),
		},
	}

	// Set upsert option to true to create document if it doesn't exist
	opts := options.Update().SetUpsert(true)

	// Always create a fresh context from Background
	updateCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := collection.UpdateOne(updateCtx, filter, update, opts)
	if err != nil {
		return false, fmt.Errorf("failed to add player ID: %v", err)
	}

	return true, nil
}
