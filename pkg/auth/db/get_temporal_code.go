package auth

import (
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetTemporalCode(code int) *authDomain.VerificationData {
	pool := config.GetPool()
	ctx := getContext()

	var result authDomain.VerificationData

	err := pool.QueryRow(ctx,
		`SELECT email, code FROM email_verification WHERE code = $1 LIMIT 1`,
		code,
	).Scan(&result.Email, &result.Code)

	if err != nil {
		return nil
	}

	return &result
}
