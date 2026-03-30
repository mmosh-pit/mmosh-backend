package bots

import (
	"context"
	"encoding/json"
	"log"
	"time"

	agentsDomain "github.com/mmosh-pit/mmosh_backend/pkg/bots/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveAgent(data *agentsDomain.CreateAgentData) error {
	pool := config.GetPool()
	ctx := context.Background()

	data.CreatedAt = time.Now()

	distributionJSON, _ := json.Marshal(data.Distribution)

	_, err := pool.Exec(ctx,
		`INSERT INTO bots (
			name, symbol, description, image, invite_image, key, lut, seniority, price,
			distribution, invitation_price, discount, telegram, twitter, website,
			presale_supply, min_presale_supply, presale_start_date, presale_end_date,
			dex_listing_date, creator, creator_username, type, code, privacy,
			default_model, created_at, status
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9,
			$10, $11, $12, $13, $14, $15,
			$16, $17, $18, $19,
			$20, $21, $22, $23, $24, $25,
			$26, $27, 'active'
		)`,
		data.Name, data.Symbol, data.Desc, data.Image, data.InviteImage, data.Key, data.Lut,
		data.Seniority, data.Price,
		distributionJSON, data.InvitationPrice, data.Discount, data.Telegram, data.Twitter,
		data.Website, data.PresaleSupply, data.MinPresaleSupply, data.PresaleStartDate,
		data.PresaleEndDate, data.DexListingDate, data.Creator, data.CreatorUsername,
		data.Type, data.Code, data.Privacy, data.DefaultModel, data.CreatedAt,
	)

	if err != nil {
		log.Printf("Could not save agent: %v\n", err)
	}

	return err
}
