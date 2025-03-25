package auth

import (
	"log"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserBySubscription(id string) (*authDomain.User, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	result := &authDomain.User{}

	err := collection.FindOne(*ctx, bson.D{{Key: "subscription.id", Value: id}}).Decode(result)

	if err == mongo.ErrNoDocuments {
		log.Printf("No document was found with the subscription %s\n", id)
		return result, common.UserNotExistsErr
	}

	return result, nil
}
