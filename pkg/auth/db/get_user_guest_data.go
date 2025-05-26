package auth

import (
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUserGuestData(userId *primitive.ObjectID) *auth.GuestUserData {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	var result auth.User

	collection := client.Database(databaseName).Collection("mmosh-users")

	res := collection.FindOne(*ctx, bson.D{{Key: "_id", Value: userId}})

	err := res.Decode(&result)

	if err != nil {
		return nil
	}

	return &result.GuestData
}
