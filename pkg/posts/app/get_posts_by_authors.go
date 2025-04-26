package posts

import (
	"fmt"
	"strings"

	postsDb "github.com/mmosh-pit/mmosh_backend/pkg/posts/db"
	postsDomain "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
)

func GetPostsByAuthors(authors []string) ([]postsDomain.Post, error) {
	if len(authors) == 0 {
		return []postsDomain.Post{}, nil
	}

	validAuthors := []string{}
	seenAuthors := make(map[string]bool)
	for _, author := range authors {
		trimmedAuthor := strings.TrimSpace(author)
		if trimmedAuthor != "" && !seenAuthors[trimmedAuthor] {
			validAuthors = append(validAuthors, trimmedAuthor)
			seenAuthors[trimmedAuthor] = true
		}
	}

	if len(validAuthors) == 0 {
		return []postsDomain.Post{}, nil
	}

	posts, err := postsDb.GetPostsByAuthors(validAuthors)
	if err != nil {
		return nil, fmt.Errorf("service error getting posts by authors: %w", err)
	}
	return posts, nil
}
