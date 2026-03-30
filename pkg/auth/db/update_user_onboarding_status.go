package auth

import (
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func UpdateUserOnboardingStatus(userId string, step int) {
	pool := config.GetPool()
	ctx := getContext()

	_, err := pool.Exec(ctx,
		`UPDATE users SET onboarding_step = $1 WHERE id = $2`,
		step, userId,
	)

	if err != nil {
		log.Printf("Could not update user onboarding step: %v\n", err)
	}
}
