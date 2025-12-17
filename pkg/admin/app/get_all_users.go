package app

import (
	adminDb "github.com/mmosh-pit/mmosh_backend/pkg/admin/db"
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
)

func GetAllUsers(page int64, search string) *[]auth.User {
	return adminDb.GetAllUsers(page, search)
}
