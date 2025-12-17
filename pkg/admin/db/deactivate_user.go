package db

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	adminDomain "github.com/mmosh-pit/mmosh_backend/pkg/admin/domain"
)

func DeactivateUser(userId *primitive.ObjectID) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	res, err := collection.UpdateOne(*ctx, bson.D{{
		Key:   "_id",
		Value: userId,
	}}, bson.D{{
		Key: "$set",
		Value: bson.D{{
			Key:   "deactivated",
			Value: true,
		}},
	}})

	if err != nil {
		return err
	}

	if res.ModifiedCount == 0 {
		return adminDomain.ErrUserNotFound
	}

	return nil
}
