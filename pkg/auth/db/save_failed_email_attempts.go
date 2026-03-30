package auth

import (
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

type FailedAttempt struct {
	Email   string
	Keypair string
}

func SaveFailedEmailAttemptsKeypairs(email, keypair string) {
	pool := config.GetPool()
	ctx := getContext()

	_, err := pool.Exec(ctx,
		`INSERT INTO failed_email_attempts (email, keypair) VALUES ($1, $2)`,
		email, keypair,
	)

	if err != nil {
		log.Printf("Error trying to save failed wallet email attempt: %v\n", err)
	}
}
