package posts

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	postsDomain "github.com/mmosh-pit/mmosh_backend/pkg/posts/domain"
	"github.com/jackc/pgx/v5/pgconn"
)

func CreatePost(post *postsDomain.Post) error {
	pool := config.GetPool()
	ctx := context.Background()

	tagsJSON, _ := json.Marshal(post.Tags)
	authorsJSON, _ := json.Marshal(post.Authors)

	err := pool.QueryRow(ctx,
		`INSERT INTO posts (header, sub_header, tags, authors, body, slug)
		 VALUES ($1, $2, $3, $4, $5, $6)
		 RETURNING id`,
		post.Header, post.SubHeader, tagsJSON, authorsJSON, post.Body, post.Slug,
	).Scan(&post.ID)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return errors.New("post with this slug already exists")
		}
		return fmt.Errorf("failed to insert post: %w", err)
	}

	return nil
}
