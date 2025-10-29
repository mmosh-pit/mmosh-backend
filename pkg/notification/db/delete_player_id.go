package notificationDb

import (
	"context"
	"fmt"
	"time"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func DeletePlayerId(playerId string, wallet string) (bool, error) {
	client, _ := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-player")

	filter := bson.M{"wallet": wallet}
	update := bson.M{
		"$pull": bson.M{"playerIds": playerId},
		"$set":  bson.M{"updated_date": time.Now().UTC()},
	}

	// Create context with timeout
	deleteCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.UpdateOne(deleteCtx, filter, update)
	if err != nil {
		return false, fmt.Errorf("failed to delete player ID: %v", err)
	}

	// Check if document was found and modified
	if result.MatchedCount == 0 {
		return false, fmt.Errorf("wallet not found: %s", wallet)
	}

	return true, nil
}
