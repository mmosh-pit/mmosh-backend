package subscriptions

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddChangePlanBadgeToSubscription(user *primitive.ObjectID) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	filter := bson.D{{Key: "_id", Value: user}}

	update := bson.D{{Key: "subscription.changed_plan", Value: true}}

	_, err := collection.UpdateOne(*ctx, filter, update)

	return err
}
