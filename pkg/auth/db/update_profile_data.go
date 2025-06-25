package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateProfileData(data auth.Profile, userId primitive.ObjectID) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	filter := bson.D{{Key: "_id", Value: userId}}

	update := bson.D{{
		Key: "$set", Value: bson.D{{
			Key:   "profile",
			Value: data,
		}},
	}}

	var newestUser auth.User

	collection.FindOne(*ctx, bson.D{{Key: "profile", Value: bson.D{{
		Key:   "$exists",
		Value: true,
	}}}}, options.FindOne().SetSort(bson.D{{Key: "seniority", Value: -1}})).Decode(&newestUser)

	data.Seniority = newestUser.Profile.Seniority

	_, err := collection.UpdateOne(*ctx, filter, update)

	return err
}
