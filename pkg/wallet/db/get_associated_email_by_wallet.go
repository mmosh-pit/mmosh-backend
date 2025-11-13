package wallet

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAssociatedEmailByWallet(wallet string) string {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-user-wallet")

	var result struct {
		Email string `json:"email"`
	}

	err := collection.FindOne(*ctx, bson.D{{
		Key: "address", Value: wallet,
	}}).Decode(&result)

	if err == mongo.ErrNoDocuments || err != nil {
		return ""
	}

	return result.Email
}
