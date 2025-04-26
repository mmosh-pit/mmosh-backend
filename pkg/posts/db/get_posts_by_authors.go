package posts

import (
	"fmt"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	posts "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetPostsByAuthors(authors []string) ([]posts.Post, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("posts")

	posts := []posts.Post{}

	filter := bson.M{"authors": bson.M{"$in": authors}}

	findOptions := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}})

	cursor, err := collection.Find(*ctx, filter, findOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to find posts by authors: %w", err)
	}
	defer cursor.Close(*ctx)

	if err = cursor.All(*ctx, &posts); err != nil {
		return nil, fmt.Errorf("failed to decode posts by authors from cursor: %w", err)
	}

	return posts, nil
}
