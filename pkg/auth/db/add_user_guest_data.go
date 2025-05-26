package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddUserGuestData(data auth.GuestUserData, userId *primitive.ObjectID) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	_, err := collection.UpdateByID(*ctx, userId, bson.D{{
		Key: "$set", Value: bson.D{{
			Key: "guest_data", Value: data,
		}},
	}})

	return err
}
