package posts

import (
	"fmt"

	postsDb "github.com/mmosh-pit/mmosh_backend/pkg/posts/db"
	postsDomain "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
)

func GetAllPosts() ([]postsDomain.Post, error) {
	posts, err := postsDb.GetAllPosts()
	if err != nil {
		return nil, fmt.Errorf("service error getting all posts: %w", err)
	}
	return posts, nil
}
