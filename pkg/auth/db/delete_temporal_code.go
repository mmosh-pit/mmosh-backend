package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func DeleteTemporalCode(code int) {
	pool := config.GetPool()
	ctx := getContext()

	pool.Exec(ctx,
		`DELETE FROM email_verification WHERE code = $1`,
		code,
	)
}
