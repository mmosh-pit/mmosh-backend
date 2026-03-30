package auth

import (
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func AddAccountDeletionRequest(data *authDomain.AccountDeletionRequest) int {
	pool := config.GetPool()
	ctx := getContext()

	var existingEmail string
	err := pool.QueryRow(ctx,
		`SELECT email FROM account_deletion_requests WHERE email = $1`,
		data.Email,
	).Scan(&existingEmail)

	if err == nil && existingEmail != "" {
		return 1
	}

	_, err = pool.Exec(ctx,
		`INSERT INTO account_deletion_requests (name, email, reason) VALUES ($1, $2, $3)`,
		data.Name, data.Email, data.Reason,
	)

	if err != nil {
		return -1
	}

	return 0
}
