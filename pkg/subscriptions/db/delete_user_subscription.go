package subscriptions

import (
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteUserSubscription(userId *primitive.ObjectID, productId string) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	log.Printf("Deleting subscription... %s", userId.String())

	collection := client.Database(databaseName).Collection("mmosh-users")

	filter := bson.D{{Key: "_id", Value: userId}, {Key: "subscription.product_id", Value: productId}}

	update := bson.D{{Key: "$unset", Value: bson.D{{Key: "subscription", Value: ""}}}}

	_, err := collection.UpdateOne(*ctx, filter, update)

	return err
}
