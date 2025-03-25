package subscriptions

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	subscriptions "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSubscriptionByProductId(productId string) (*subscriptions.Subscription, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("subscription")

	var subscription subscriptions.Subscription

	err := collection.FindOne(*ctx, bson.D{{Key: "product_id", Value: productId}}).Decode(&subscription)

	return &subscription, err
}
