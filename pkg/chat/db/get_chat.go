package chat

import (
	"context"
	"errors"
	"log"

	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/jackc/pgx/v5"
)

func GetChat(userId string) (*chatDomain.Chat, error) {
	pool := config.GetPool()
	ctx := context.Background()

	row := pool.QueryRow(ctx,
		`SELECT `+chatSelectColumns+` FROM chats WHERE owner = $1 LIMIT 1`,
		userId,
	)

	result, err := scanChat(row)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, common.ChatNotExistsErr
	}

	if err != nil {
		log.Printf("Got error: %v\n", err)
		return nil, err
	}

	return &result, nil
}
