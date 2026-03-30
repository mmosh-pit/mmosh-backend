package db

import (
	"context"
	"log"
	"time"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/jackc/pgx/v5"
)

func GetAllUsers(page int64, search string) *[]auth.User {
	pool := config.GetPool()
	ctx := context.Background()

	var (
		rows pgx.Rows
		err  error
	)

	if search != "" {
		pattern := "%" + search + "%"
		rows, err = pool.Query(ctx,
			`SELECT id, name, email, last_login, created_at, deactivated, role
			 FROM users
			 WHERE name ILIKE $1 OR username ILIKE $1 OR email ILIKE $1
			 ORDER BY seniority DESC
			 LIMIT 20 OFFSET $2`,
			pattern, page*20,
		)
	} else {
		rows, err = pool.Query(ctx,
			`SELECT id, name, email, last_login, created_at, deactivated, role
			 FROM users
			 ORDER BY seniority DESC
			 LIMIT 20 OFFSET $1`,
			page*20,
		)
	}

	var result []auth.User

	if err != nil {
		log.Printf("[ADMIN/GET ALL USERS] Got error retrieving all users: %v\n", err)
		return &result
	}
	defer rows.Close()

	for rows.Next() {
		var u auth.User
		var lastLogin *time.Time

		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &lastLogin, &u.CreatedAt, &u.Deactivated, &u.Role); err != nil {
			log.Printf("[ADMIN/GET ALL USERS] could not decode user: %v\n", err)
			continue
		}

		if lastLogin != nil {
			u.LastLogin = *lastLogin
		}

		result = append(result, u)
	}

	return &result
}
