package auth

import (
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

type FailedAttempt struct {
	Email   string `bson:"email"`
	Keypair string `bson:"keypair"`
}

func SaveFailedEmailAttemptsKeypairs(email, keypair string) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-user-wallet")

	data := FailedAttempt{
		Email:   email,
		Keypair: keypair,
	}

	_, err := collection.InsertOne(*ctx, data)

	if err != nil {
		log.Printf("Error trying to save failed wallet email attempt: %v\n", err)
	}
}
