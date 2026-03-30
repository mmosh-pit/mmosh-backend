package auth

import (
	"encoding/json"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func UpdateProfileData(data auth.User, userId string) error {
	pool := config.GetPool()
	ctx := getContext()

	var maxSeniority int
	pool.QueryRow(ctx, `SELECT COALESCE(MAX(seniority), 0) FROM users WHERE profilenft IS NOT NULL`).Scan(&maxSeniority)
	data.Seniority = maxSeniority

	websitesJSON, _ := json.Marshal(data.Websites)

	_, err := pool.Exec(ctx,
		`UPDATE users SET name = $1, display_name = $2, username = $3, bio = $4,
		  picture = $5, banner = $6, websites = $7, symbol = $8, link = $9,
		  challenges = $10, seniority = $11
		 WHERE id = $12`,
		data.Name, data.DisplayName, data.Username, data.Bio,
		data.Picture, data.Banner, websitesJSON, data.Symbol, data.Link,
		data.Challenges, data.Seniority, userId,
	)

	return err
}
