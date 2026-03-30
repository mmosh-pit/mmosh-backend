package posts

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	postsDomain "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
)

func GetAllPosts() ([]postsDomain.Post, error) {
	pool := config.GetPool()
	ctx := context.Background()

	rows, err := pool.Query(ctx,
		`SELECT id, header, sub_header, tags, authors, body, slug, created_at, updated_at
		 FROM posts ORDER BY created_at DESC`,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to find posts: %w", err)
	}
	defer rows.Close()

	var posts []postsDomain.Post

	for rows.Next() {
		var p postsDomain.Post
		var tagsJSON, authorsJSON []byte

		if err := rows.Scan(&p.ID, &p.Header, &p.SubHeader, &tagsJSON, &authorsJSON, &p.Body, &p.Slug, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, fmt.Errorf("failed to decode posts from cursor: %w", err)
		}

		if len(tagsJSON) > 0 {
			json.Unmarshal(tagsJSON, &p.Tags)
		}
		if len(authorsJSON) > 0 {
			json.Unmarshal(authorsJSON, &p.Authors)
		}

		posts = append(posts, p)
	}

	if posts == nil {
		posts = []postsDomain.Post{}
	}

	return posts, nil
}
