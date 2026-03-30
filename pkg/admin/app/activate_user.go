package app

import (
	adminDb "github.com/mmosh-pit/mmosh_backend/pkg/admin/db"
)

func ActivateUser(userId string) error {
	return adminDb.ActivateUser(userId)
}
