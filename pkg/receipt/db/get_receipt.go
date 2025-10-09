package receiptDb

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	receiptDomain "github.com/mmosh-pit/mmosh_backend/pkg/receipt/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetReceipt retrieves a receipt by purchase token from MongoDB
func GetReceipt(purchaseToken string) (*receiptDomain.Receipt, error) {
	client, ctx := config.GetMongoClient()
	dbName := config.GetDatabaseName()

	collection := client.Database(dbName).Collection("mmosh-app-receipt")

	filter := bson.M{"purchase_token": purchaseToken}

	var receipt receiptDomain.Receipt
	err := collection.FindOne(*ctx, filter).Decode(&receipt)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &receipt, nil
}
