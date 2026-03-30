package auth

import (
	"encoding/json"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func CreateUser(data *authDomain.User) error {
	pool := config.GetPool()
	ctx := getContext()

	sessionsJSON, _ := json.Marshal(data.Sessions)

	err := pool.QueryRow(ctx,
		`INSERT INTO users (uuid, email, password, name, picture, sessions, wallet, referred_by,
		  onboarding_step, created_at, role, from_bot)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		 RETURNING id`,
		data.UUID, data.Email, data.Password, data.Name, data.Picture,
		sessionsJSON, data.Wallet, data.ReferredBy, data.OnboardingStep,
		data.CreatedAt, data.Role, data.FromBot,
	).Scan(&data.ID)

	return err
}
