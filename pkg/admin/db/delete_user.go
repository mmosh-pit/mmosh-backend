package db

import (
	"context"

	adminDomain "github.com/mmosh-pit/mmosh_backend/pkg/admin/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func DeleteUser(userId string) error {
	pool := config.GetPool()
	ctx := context.Background()

	tag, err := pool.Exec(ctx,
		`DELETE FROM users WHERE id = $1`,
		userId,
	)

	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return adminDomain.ErrUserNotFound
	}

	return nil
}
