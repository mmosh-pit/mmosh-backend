package posts

import (
	"errors"
	"fmt"
	"strings"

	postsDb "github.com/mmosh-pit/mmosh_backend/pkg/posts/db"
	postsDomain "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
)

func GetPostBySlug(slug string) (*postsDomain.Post, error) {
	if strings.TrimSpace(slug) == "" {
		return nil, errors.New("slug cannot be empty")
	}

	post, err := postsDb.GetPostBySlug(slug)
	if err != nil {
		return nil, fmt.Errorf("service error getting post by slug: %w", err)
	}
	return post, nil
}
