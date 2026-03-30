package posts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	postsDomain "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
)

func GetPostBySlug(slug string) (*postsDomain.Post, error) {
	pool := config.GetPool()
	ctx := context.Background()

	var p postsDomain.Post
	var tagsJSON, authorsJSON []byte

	err := pool.QueryRow(ctx,
		`SELECT id, header, sub_header, tags, authors, body, slug, created_at, updated_at
		 FROM posts WHERE slug = $1`,
		slug,
	).Scan(&p.ID, &p.Header, &p.SubHeader, &tagsJSON, &authorsJSON, &p.Body, &p.Slug, &p.CreatedAt, &p.UpdatedAt)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, errors.New("post not found")
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find post by slug '%s': %w", slug, err)
	}

	if len(tagsJSON) > 0 {
		json.Unmarshal(tagsJSON, &p.Tags)
	}
	if len(authorsJSON) > 0 {
		json.Unmarshal(authorsJSON, &p.Authors)
	}

	return &p, nil
}
