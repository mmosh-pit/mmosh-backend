package members

import (
	"context"
	"log"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/jackc/pgx/v5"
)

func GetMembers(page int64, search string, userId string) []auth.User {
	pool := config.GetPool()
	ctx := context.Background()

	result := []auth.User{}

	var (
		rows pgx.Rows
		err  error
	)

	if search != "" {
		pattern := "%" + search + "%"
		rows, err = pool.Query(ctx,
			`SELECT id, name, username, picture, email, seniority
			 FROM users
			 WHERE id != $1
			   AND profilenft IS NOT NULL
			   AND (name ILIKE $2 OR username ILIKE $2)
			 ORDER BY seniority DESC
			 LIMIT 20 OFFSET $3`,
			userId, pattern, page*20,
		)
	} else {
		rows, err = pool.Query(ctx,
			`SELECT id, name, username, picture, email, seniority
			 FROM users
			 WHERE id != $1
			   AND profilenft IS NOT NULL
			 ORDER BY seniority DESC
			 LIMIT 20 OFFSET $2`,
			userId, page*20,
		)
	}

	if err != nil {
		log.Printf("Got error retrieving members: %v\n", err)
		return result
	}
	defer rows.Close()

	for rows.Next() {
		var u auth.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Username, &u.Picture, &u.Email, &u.Seniority); err != nil {
			log.Printf("Error decoding user: %v\n", err)
			continue
		}
		result = append(result, u)
	}

	return result
}
