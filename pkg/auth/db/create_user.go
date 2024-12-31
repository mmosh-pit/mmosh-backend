package auth

import (
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(data *authDomain.User) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-users")

	res, err := collection.InsertOne(*ctx, *data)

	if err != nil {
		return err
	}

	id := res.InsertedID.(primitive.ObjectID)

	data.ID = &id

	return nil
}
