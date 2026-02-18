package auth

import (
	"time"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SaveLastLogin(id *primitive.ObjectID) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	filter := bson.D{{Key: "_id", Value: id}}

	update := bson.D{{
		Key: "$set", Value: bson.D{{
			Key:   "last_login",
			Value: time.Now(),
		}},
	}}

	collection.UpdateOne(*ctx, filter, update)
}
