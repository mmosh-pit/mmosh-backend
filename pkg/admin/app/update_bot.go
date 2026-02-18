package app

import (
	adminDb "github.com/mmosh-pit/mmosh_backend/pkg/admin/db"
	adminDomain "github.com/mmosh-pit/mmosh_backend/pkg/admin/domain"
)

func UpdateBot(payload adminDomain.UpdateBotPayload) {
	adminDb.UpdateBot(payload)
}
