package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetOnboardingStep(userId string, step int) {

	userIdBson, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		return
	}

	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	_, err = collection.UpdateByID(*ctx, userIdBson, bson.D{{
		Key: "$set", Value: bson.D{{
			Key: "onboarding_step", Value: step,
		}},
	}})
}
