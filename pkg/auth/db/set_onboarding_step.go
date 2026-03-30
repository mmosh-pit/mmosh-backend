package auth

import (
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SetOnboardingStep(userId string, step int) {
	pool := config.GetPool()
	ctx := getContext()

	pool.Exec(ctx,
		`UPDATE users SET onboarding_step = $1 WHERE id = $2`,
		step, userId,
	)
}
