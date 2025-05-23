package subscriptions

import (
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	subscriptionsDomain "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/domain"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserSubscription(userId *primitive.ObjectID, productId string, expiresAt int64, platform string) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	database := client.Database(databaseName)

	userCollection := database.Collection("mmosh-users")

	subscriptionsCollection := database.Collection("subscription")

	var subscriptionToAdd subscriptionsDomain.Subscription

	err := subscriptionsCollection.FindOne(*ctx, bson.D{{Key: "product_id", Value: productId}}).Decode(&subscriptionToAdd)

	if err != nil {
		return err
	}

	userSubscription := authDomain.UserSubscription{
		ProductId:        productId,
		PurchaseToken:    "",
		SubscriptionId:   subscriptionToAdd.ID.String(),
		SubscriptionTier: subscriptionToAdd.Tier,
		ExpiresAt:        expiresAt,
		Platform:         platform,
	}

	_, err = userCollection.UpdateByID(*ctx, userId, bson.D{{Key: "$set", Value: bson.D{{Key: "subscription", Value: userSubscription}}}})

	return err
}
