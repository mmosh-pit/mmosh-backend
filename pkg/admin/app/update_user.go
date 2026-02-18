package app

import (
	adminDb "github.com/mmosh-pit/mmosh_backend/pkg/admin/db"
	adminDomain "github.com/mmosh-pit/mmosh_backend/pkg/admin/domain"
)

func UpdateUser(payload adminDomain.UpdateUserPayload) {
	adminDb.UpdateUser(payload)
}
