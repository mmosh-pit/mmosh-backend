package auth

import (
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func GetWalletByEmail(email string) *auth.Wallet {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-user-wallet")

	var data auth.Wallet

	err := collection.FindOne(*ctx, bson.D{{Key: "email", Value: email}}).Decode(&data)

	if err != nil {
		log.Printf("Error querying wallet: %v\n", err)

		return nil
	}

	return &data
}
