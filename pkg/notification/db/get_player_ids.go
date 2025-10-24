package notificationDb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

type PlayerData struct {
	Wallet    string   `bson:"wallet"`
	PlayerIDs []string `bson:"playerIds"`
}

func GetPlayerIds(wallet string) ([]string, error) {
	client, _ := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-player")

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Find the wallet document
	filter := bson.M{"wallet": wallet}
	var userData PlayerData

	err := collection.FindOne(ctx, filter).Decode(&userData)
	if err != nil {
		log.Printf("Error finding wallet: %v", err)
		return nil, fmt.Errorf("wallet not found")
	}

	// Check if playerIds exist
	if len(userData.PlayerIDs) == 0 {
		return nil, fmt.Errorf("no player IDs found for this wallet")
	}

	return userData.PlayerIDs, nil
}
