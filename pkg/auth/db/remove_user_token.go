package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RemoveUserToken(userId *primitive.ObjectID, token string) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	filter := bson.D{{Key: "_id", Value: userId}}

	update := bson.D{{Key: "$pull", Value: bson.D{{Key: "sessions", Value: token}}}}

	_, err := collection.UpdateOne(*ctx, filter, update)

	return err
}
