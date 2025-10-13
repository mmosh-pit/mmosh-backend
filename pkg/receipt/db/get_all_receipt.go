package receiptDb

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	receiptDomain "github.com/mmosh-pit/mmosh_backend/pkg/receipt/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func GetAllReceipts() ([]receiptDomain.Receipt, error) {
	client, ctx := config.GetMongoClient()
	dbName := config.GetDatabaseName()

	collection := client.Database(dbName).Collection("mmosh-app-receipt")

	cursor, err := collection.Find(*ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(*ctx)

	var receipts []receiptDomain.Receipt
	for cursor.Next(*ctx) {
		var receipt receiptDomain.Receipt
		if err := cursor.Decode(&receipt); err != nil {
			return nil, err
		}
		receipts = append(receipts, receipt)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return receipts, nil
}
