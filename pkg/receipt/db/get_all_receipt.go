package receiptDb

import (
	"context"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	receiptDomain "github.com/mmosh-pit/mmosh_backend/pkg/receipt/domain"
)

func GetAllReceipts() ([]receiptDomain.Receipt, error) {
	pool := config.GetPool()
	ctx := context.Background()

	rows, err := pool.Query(ctx,
		`SELECT id, package_name, product_id, purchase_token, wallet, platform, created_at, expired_at, is_canceled
		 FROM receipts`,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var receipts []receiptDomain.Receipt

	for rows.Next() {
		var r receiptDomain.Receipt
		if err := rows.Scan(
			&r.ID, &r.PackageName, &r.ProductID, &r.PurchaseToken, &r.Wallet,
			&r.Platform, &r.CreatedAt, &r.ExpiredAt, &r.IsCanceled,
		); err != nil {
			return nil, err
		}
		receipts = append(receipts, r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return receipts, nil
}
