package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveTelegramData(data auth.TelegramUserData, userId primitive.ObjectID) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	filter := bson.D{{Key: "_id", Value: userId}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "telegram", Value: data}}}}

	_, err := collection.UpdateOne(
		*ctx, filter, update,
	)

	return err
}
