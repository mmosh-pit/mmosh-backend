package receiptDb

import (
	"context"
	"time"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	receiptDomain "github.com/mmosh-pit/mmosh_backend/pkg/receipt/domain"
)

func SaveReceipt(data *receiptDomain.Receipt) error {
	pool := config.GetPool()
	ctx := context.Background()

	err := pool.QueryRow(ctx,
		`INSERT INTO receipts (package_name, product_id, purchase_token, wallet, platform, created_at, expired_at, is_canceled)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		 RETURNING id`,
		data.PackageName, data.ProductID, data.PurchaseToken, data.Wallet, data.Platform,
		time.Now(), data.ExpiredAt, data.IsCanceled,
	).Scan(&data.ID)

	return err
}
