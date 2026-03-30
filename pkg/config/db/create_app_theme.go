package db

import (
	"context"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	configDomain "github.com/mmosh-pit/mmosh_backend/pkg/config/domain"
)

func CreateAppTheme(data configDomain.AppTheme) {
	pool := config.GetPool()
	ctx := context.Background()

	pool.Exec(ctx,
		`INSERT INTO themes (name, code_name, background_color, primary_color, secondary_color, logo)
		 VALUES ($1, $2, $3, $4, $5, $6)`,
		data.Name, data.CodeName, data.BackgroundColor, data.PrimaryColor, data.SecondaryColor, data.Logo,
	)
}
