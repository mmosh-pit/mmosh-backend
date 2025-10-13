package receiptDb

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteReceipt(purchaseToken string) error {
	client, ctx := config.GetMongoClient()
	dbName := config.GetDatabaseName()

	collection := client.Database(dbName).Collection("mmosh-app-receipt")

	filter := bson.M{"purchase_token": purchaseToken}

	_, err := collection.DeleteOne(*ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
