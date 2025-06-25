package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteBlueskyData(userId primitive.ObjectID) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	filter := bson.D{{Key: "_id", Value: userId}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "bluesky", Value: nil}}}}

	_, err := collection.UpdateOne(
		*ctx, filter, update,
	)

	return err
}
