package auth

import (
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUserBySessionToken(token string) (*authDomain.User, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	var user authDomain.User

	err := collection.FindOne(*ctx, bson.D{{Key: "sessions", Value: token}}).Decode(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
