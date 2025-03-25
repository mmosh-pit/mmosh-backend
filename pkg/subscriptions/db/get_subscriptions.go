package subscriptions

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"
	"go.mongodb.org/mongo-driver/bson"
)

func GetSubscriptions() ([]subscriptionsDomain.Subscription, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("subscription")

	var subscriptions []subscriptionsDomain.Subscription

	res, err := collection.Find(*ctx, bson.D{{}})

	for res.Next(*ctx) {
		var subscription subscriptionsDomain.Subscription
		err = res.Decode(&subscription)

		if err != nil {
			return subscriptions, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, err
}
