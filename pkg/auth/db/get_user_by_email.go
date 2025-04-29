package auth

import (
	"log"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserByEmail(email string) (authDomain.User, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	var result authDomain.User

	err := collection.FindOne(*ctx, bson.D{{Key: "email", Value: email}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		log.Printf("No document was found with the title %s\n", email)
		return result, err
	}

	result.Password = ""

	return result, nil
}
