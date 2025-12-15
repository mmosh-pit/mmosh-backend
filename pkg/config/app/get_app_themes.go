package app

import (
	configDb "github.com/mmosh-pit/mmosh_backend/pkg/config/db"
	configDomain "github.com/mmosh-pit/mmosh_backend/pkg/config/domain"
)

func GetAppThemes() *[]configDomain.AppTheme {
	return configDb.GetAppThemes()
}
