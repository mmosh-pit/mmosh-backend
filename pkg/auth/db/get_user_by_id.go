package auth

import (
	"log"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserById(id string) (authDomain.User, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	var result authDomain.User

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Printf("Invalid Object ID %s\n", id)
		return result, err
	}

	err = collection.FindOne(*ctx, bson.D{{Key: "_id", Value: objectId}}).Decode(&result)

	if err == mongo.ErrNoDocuments {
		log.Printf("No document was found with the title %s\n", id)
		return result, err
	}

	result.Password = ""

	return result, nil
}
