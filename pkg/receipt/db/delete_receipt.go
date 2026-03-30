package receiptDb

import (
	"context"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func DeleteReceipt(purchaseToken string) error {
	pool := config.GetPool()
	ctx := context.Background()

	_, err := pool.Exec(ctx,
		`DELETE FROM receipts WHERE purchase_token = $1`,
		purchaseToken,
	)

	return err
}
