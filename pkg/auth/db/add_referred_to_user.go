package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddReferrerToUser(username string, userId *primitive.ObjectID) error {

	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	_, err := collection.UpdateByID(*ctx, userId, bson.D{{
		Key: "$set", Value: bson.D{{
			Key:   "referred_by",
			Value: username,
		}},
	}})

	return err
}
