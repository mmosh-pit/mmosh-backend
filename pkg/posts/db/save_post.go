package posts

import (
	"errors"
	"fmt"
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	posts "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreatePost(post *posts.Post) error {
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("posts")

	post.ID = primitive.NilObjectID
	res, err := collection.InsertOne(*ctx, post)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.New("post with this slug already exists")
		}
		return fmt.Errorf("failed to insert post: %w", err)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		post.ID = oid
	} else {
		log.Printf("WARN: Could not cast InsertedID to ObjectID for slug: %s", post.Slug)
	}
	return nil
}
