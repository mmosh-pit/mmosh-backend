package receiptDb

import (
	"time"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

// UpdateReceipt updates the ExpiredAt field of a receipt by purchase token
func UpdateReceipt(purchaseToken string, isCanceled bool) error {
	client, ctx := config.GetMongoClient()
	dbName := config.GetDatabaseName()

	collection := client.Database(dbName).Collection("mmosh-app-receipt")

	filter := bson.M{"purchase_token": purchaseToken}

	updateFields := bson.M{}

	if isCanceled {
		updateFields["is_canceled"] = true
	} else {
		// "expired_at": time.Now().UTC().Add(31 * 24 * time.Hour),
		updateFields["expired_at"] = time.Now().UTC().Add(5 * time.Minute)
	}

	update := bson.M{"$set": updateFields}

	_, err := collection.UpdateOne(*ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}
