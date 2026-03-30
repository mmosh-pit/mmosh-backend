package receiptDb

import (
	"context"
	"time"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func UpdateReceipt(purchaseToken string, isCanceled bool) error {
	pool := config.GetPool()
	ctx := context.Background()

	var err error

	if isCanceled {
		_, err = pool.Exec(ctx,
			`UPDATE receipts SET is_canceled = true WHERE purchase_token = $1`,
			purchaseToken,
		)
	} else {
		_, err = pool.Exec(ctx,
			`UPDATE receipts SET expired_at = $1 WHERE purchase_token = $2`,
			time.Now().UTC().Add(5*time.Minute), purchaseToken,
		)
	}

	return err
}
