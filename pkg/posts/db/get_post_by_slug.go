package posts

import (
	"errors"
	"fmt"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	postsDomain "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetPostBySlug(slug string) (*postsDomain.Post, error) {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("posts")

	var post postsDomain.Post
	filter := bson.M{"slug": slug}

	err := collection.FindOne(*ctx, filter).Decode(&post)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("post not found")
		}
		return nil, fmt.Errorf("failed to find post by slug '%s': %w", slug, err)
	}
	return &post, nil
}
