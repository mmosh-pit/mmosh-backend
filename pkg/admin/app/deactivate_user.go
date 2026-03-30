package app

import (
	adminDb "github.com/mmosh-pit/mmosh_backend/pkg/admin/db"
)

func DeactivateUser(userId string) error {
	return adminDb.DeactivateUser(userId)
}
