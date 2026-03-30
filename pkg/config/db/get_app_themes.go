package db

import (
	"context"
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	configDomain "github.com/mmosh-pit/mmosh_backend/pkg/config/domain"
)

func GetAppThemes() *[]configDomain.AppTheme {
	pool := config.GetPool()
	ctx := context.Background()

	var result []configDomain.AppTheme

	rows, err := pool.Query(ctx,
		`SELECT id, name, code_name, background_color, primary_color, secondary_color, logo FROM themes`,
	)

	if err != nil {
		log.Printf("[APP THEMES] could not get themes: %v\n", err)
		return &result
	}
	defer rows.Close()

	for rows.Next() {
		var theme configDomain.AppTheme
		if err := rows.Scan(
			&theme.ID, &theme.Name, &theme.CodeName,
			&theme.BackgroundColor, &theme.PrimaryColor, &theme.SecondaryColor, &theme.Logo,
		); err != nil {
			log.Printf("[APP THEMES] Got error decoding theme: %v\n", err)
			continue
		}
		result = append(result, theme)
	}

	return &result
}
