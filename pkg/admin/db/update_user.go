package db

import (
	"context"
	"log"

	adminDomain "github.com/mmosh-pit/mmosh_backend/pkg/admin/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func UpdateUser(payload adminDomain.UpdateUserPayload) {
	pool := config.GetPool()
	ctx := context.Background()

	_, err := pool.Exec(ctx,
		`UPDATE users SET name = $1, email = $2, role = $3, deactivated = $4 WHERE id = $5`,
		payload.Name, payload.Email, payload.Role, payload.Deactivated, payload.ID,
	)

	if err != nil {
		log.Printf("Error updating USER: %v\n", err)
	}
}
