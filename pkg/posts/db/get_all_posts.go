package posts

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	postsDomain "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
)

func GetAllPosts() ([]postsDomain.Post, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("posts")

	var posts []postsDomain.Post

	findOptions := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}})

	cursor, err := collection.Find(*ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to find posts: %w", err)
	}

	defer cursor.Close(*ctx)

	if err = cursor.All(*ctx, &posts); err != nil {
		return nil, fmt.Errorf("failed to decode posts from cursor: %w", err)
	}

	if posts == nil {
		posts = []postsDomain.Post{}
	}

	return posts, nil
}
