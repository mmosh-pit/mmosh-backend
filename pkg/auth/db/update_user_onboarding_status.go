package auth

import (
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateUserOnboardingStatus(userId *primitive.ObjectID, step int) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	_, err := collection.UpdateByID(*ctx, userId, bson.D{{
		Key: "$set", Value: bson.D{{
			Key:   "onboarding_step",
			Value: step,
		}},
	}})

	if err != nil {
		log.Printf("Could not update user onboarding step: %v\n", err)
	}
}
