package app

import (
	adminDb "github.com/mmosh-pit/mmosh_backend/pkg/admin/db"
	bots "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
)

func GetAllBots(page int, search string) *[]bots.Bot {
	return adminDb.GetAllBots(int64(page), search)
}
