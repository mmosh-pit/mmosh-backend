package receiptDb

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	receiptDomain "github.com/mmosh-pit/mmosh_backend/pkg/receipt/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveReceipt(data *receiptDomain.Receipt) error {
	client, ctx := config.GetMongoClient()
	dbName := config.GetDatabaseName()

	collection := client.Database(dbName).Collection("mmosh-app-receipt")

	res, err := collection.InsertOne(*ctx, *data)
	if err != nil {
		return err
	}

	id := res.InsertedID.(primitive.ObjectID)
	data.ID = id

	return nil
}
