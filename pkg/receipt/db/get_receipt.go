package receiptDb

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	receiptDomain "github.com/mmosh-pit/mmosh_backend/pkg/receipt/domain"
)

func GetReceipt(purchaseToken string) (*receiptDomain.Receipt, error) {
	pool := config.GetPool()
	ctx := context.Background()

	var r receiptDomain.Receipt

	err := pool.QueryRow(ctx,
		`SELECT id, package_name, product_id, purchase_token, wallet, platform, created_at, expired_at, is_canceled
		 FROM receipts WHERE purchase_token = $1`,
		purchaseToken,
	).Scan(
		&r.ID, &r.PackageName, &r.ProductID, &r.PurchaseToken, &r.Wallet,
		&r.Platform, &r.CreatedAt, &r.ExpiredAt, &r.IsCanceled,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &r, nil
}
