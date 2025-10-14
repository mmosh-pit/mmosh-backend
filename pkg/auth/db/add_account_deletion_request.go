package auth

import (
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func AddAccountDeletionRequest(data *authDomain.AccountDeletionRequest) int {

	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("account-deletion-requests")

	var existingData authDomain.AccountDeletionRequest

	_ = collection.FindOne(*ctx, bson.D{{Key: "email", Value: data.Email}}).Decode(&existingData)

	if existingData.Email != "" {
		return 1
	}

	_, err := collection.InsertOne(*ctx, data)

	if err != nil {
		return -1
	}

	return 0
}
