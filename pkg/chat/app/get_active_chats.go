package chat

import (
	"errors"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	chatDb "github.com/mmosh-pit/mmosh_backend/pkg/chat/db"
	chatDomain "github.com/mmosh-pit/mmosh_backend/pkg/chat/domain"
	commonDomain "github.com/mmosh-pit/mmosh_backend/pkg/common/domain"
	"github.com/jackc/pgx/v5"
)

func GetActiveChats(ownerId string) ([]chatDomain.Chat, error) {
	_, err := auth.GetUserById(ownerId)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, commonDomain.UserNotExistsErr
		}

		return nil, err
	}

	chats := chatDb.GetActiveChats(ownerId)

	return chats, nil
}
