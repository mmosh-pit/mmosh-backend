package wallet

import (
	"context"
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func GetAllCoinAddresses() ([]string, error) {
	pool := config.GetPool()
	ctx := context.Background()

	rows, err := pool.Query(ctx, `SELECT token FROM coin_addresses`)

	var resultingTokens []string

	if err != nil {
		return resultingTokens, err
	}
	defer rows.Close()

	for rows.Next() {
		var token string
		if err := rows.Scan(&token); err != nil {
			log.Fatal(err)
		}
		resultingTokens = append(resultingTokens, token)
	}

	return resultingTokens, nil
}
