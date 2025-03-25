package subscriptions

import (
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateSubscriptionStatus(user *primitive.ObjectID, productId string, newExpiresAt int64) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	subscriptions, _ := GetSubscriptions()

	collection := client.Database(databaseName).Collection("mmosh-users")

	var subscription subscriptionsDomain.Subscription

	for _, value := range subscriptions {
		if value.ProductId == productId {
			subscription = value
		}
	}

	filter := bson.D{{Key: "_id", Value: user}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "subscription.expires_at", Value: newExpiresAt}, {Key: "subscription.subscription_tier", Value: subscription.Tier}, {Key: "subscription.product_id", Value: productId}, {Key: "subscription.subscription_id", Value: subscription.ID}, {Key: "subscription.changed_plan", Value: false}}}}

	_, err := collection.UpdateOne(*ctx, filter, update)

	return err
}
