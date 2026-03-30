package auth

import (
	"context"
	"encoding/json"
	"time"

	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/jackc/pgx/v5"
)

const selectUserColumns = `id, uuid, picture, banner, name, display_name, last_name, username,
    websites, bio, challenges, email, password, telegram, sessions, bluesky, subscription,
    wallet, referred_by, onboarding_step, created_at, last_login, profilenft, role, from_bot,
    deactivated, seniority, symbol, link, following, follower, connection_nft, connection_badge,
    connection, is_private, request`

func scanUser(row pgx.Row) (authDomain.User, error) {
	var u authDomain.User
	var websitesJSON, telegramJSON, sessionsJSON, blueskyJSON, subscriptionJSON []byte
	var lastLogin *time.Time

	err := row.Scan(
		&u.ID, &u.UUID, &u.Picture, &u.Banner, &u.Name, &u.DisplayName, &u.LastName,
		&u.Username, &websitesJSON, &u.Bio, &u.Challenges, &u.Email, &u.Password,
		&telegramJSON, &sessionsJSON, &blueskyJSON, &subscriptionJSON,
		&u.Wallet, &u.ReferredBy, &u.OnboardingStep, &u.CreatedAt, &lastLogin,
		&u.ProfileNFT, &u.Role, &u.FromBot, &u.Deactivated,
		&u.Seniority, &u.Symbol, &u.Link, &u.Following, &u.Follower,
		&u.ConnectionNFT, &u.ConnectionBadge, &u.Connection, &u.IsPrivate, &u.Request,
	)

	if err != nil {
		return u, err
	}

	if lastLogin != nil {
		u.LastLogin = *lastLogin
	}

	unmarshalUserJSON(&u, websitesJSON, telegramJSON, sessionsJSON, blueskyJSON, subscriptionJSON)

	return u, nil
}

func scanUserRows(rows pgx.Rows) ([]authDomain.User, error) {
	var users []authDomain.User

	for rows.Next() {
		var u authDomain.User
		var websitesJSON, telegramJSON, sessionsJSON, blueskyJSON, subscriptionJSON []byte
		var lastLogin *time.Time

		err := rows.Scan(
			&u.ID, &u.UUID, &u.Picture, &u.Banner, &u.Name, &u.DisplayName, &u.LastName,
			&u.Username, &websitesJSON, &u.Bio, &u.Challenges, &u.Email, &u.Password,
			&telegramJSON, &sessionsJSON, &blueskyJSON, &subscriptionJSON,
			&u.Wallet, &u.ReferredBy, &u.OnboardingStep, &u.CreatedAt, &lastLogin,
			&u.ProfileNFT, &u.Role, &u.FromBot, &u.Deactivated,
			&u.Seniority, &u.Symbol, &u.Link, &u.Following, &u.Follower,
			&u.ConnectionNFT, &u.ConnectionBadge, &u.Connection, &u.IsPrivate, &u.Request,
		)

		if err != nil {
			continue
		}

		if lastLogin != nil {
			u.LastLogin = *lastLogin
		}

		unmarshalUserJSON(&u, websitesJSON, telegramJSON, sessionsJSON, blueskyJSON, subscriptionJSON)
		users = append(users, u)
	}

	return users, rows.Err()
}

func unmarshalUserJSON(u *authDomain.User, websites, telegram, sessions, bluesky, subscription []byte) {
	if len(sessions) > 0 {
		json.Unmarshal(sessions, &u.Sessions)
	}
	if len(telegram) > 0 {
		json.Unmarshal(telegram, &u.Telegram)
	}
	if len(bluesky) > 0 {
		json.Unmarshal(bluesky, &u.Bluesky)
	}
	if len(subscription) > 0 {
		json.Unmarshal(subscription, &u.Subscription)
	}
	if len(websites) > 0 {
		json.Unmarshal(websites, &u.Websites)
	}
}

func getContext() context.Context {
	return context.Background()
}
