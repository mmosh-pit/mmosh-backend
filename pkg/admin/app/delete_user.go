package app

import (
	adminDb "github.com/mmosh-pit/mmosh_backend/pkg/admin/db"
)

func DeleteUser(userId string) error {
	return adminDb.DeleteUser(userId)
}
