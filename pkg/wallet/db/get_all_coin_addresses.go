package wallet

import (
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TokenOnly struct {
	Token string `bson:"token"`
}

func GetAllCoinAddresses() ([]string, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-tokens")

	res, err := collection.Find(*ctx, bson.D{}, options.Find().SetProjection(bson.D{{Key: "token", Value: 1}}))

	var resultingTokens []string

	if err != nil {

		if err == mongo.ErrNoDocuments {
			return resultingTokens, nil
		}

		return resultingTokens, err
	}

	for res.Next(*ctx) {
		var token TokenOnly

		if err := res.Decode(&token); err != nil {
			log.Fatal(err)
		}
		resultingTokens = append(resultingTokens, token.Token)
	}

	return resultingTokens, nil
}
