package receiptApp

import (
	"context"
	"fmt"
	"time"

	receiptDb "github.com/mmosh-pit/mmosh_backend/pkg/receipt/db"
	receiptDomain "github.com/mmosh-pit/mmosh_backend/pkg/receipt/domain"
	"google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
)

func VerifyReceipt(params *receiptDomain.VerifyReceiptParams, authToken string) (*receiptDomain.PaymentStatus, error) {
	packageInfo, err := receiptDb.GetReceipt(params.PurchaseToken)
	if err != nil {
		return nil, fmt.Errorf("error fetching receipt from database: %v", err)
	}

	if packageInfo == nil {
		return nil, fmt.Errorf("receipt not found in database")
	}

	paymentInfo, err := CheckPaymentStatus(packageInfo.PackageName, packageInfo.ProductID, params.PurchaseToken)
	if err != nil {
		return nil, fmt.Errorf("error checking payment status: %v", err)
	}
	return &paymentInfo.Status, nil
}

func CheckPaymentStatus(packageName, productId, token string) (*receiptDomain.PaymentInfo, error) {
	ctx := context.Background()

	srv, err := androidpublisher.NewService(ctx, option.WithCredentialsFile("service-account.json"))
	if err != nil {
		return nil, fmt.Errorf("error creating android publisher service: %v", err)
	}

	purchase, err := srv.Purchases.Subscriptions.Get(packageName, productId, token).Do()
	if err != nil {
		return nil, fmt.Errorf("error fetching subscription: %v", err)
	}

	paymentInfo := &receiptDomain.PaymentInfo{
		AutoRenewing: purchase.AutoRenewing,
	}

	if purchase.ExpiryTimeMillis > 0 {
		paymentInfo.ExpiryTime = time.Unix(0, purchase.ExpiryTimeMillis*int64(time.Millisecond))
	}
	if purchase.StartTimeMillis > 0 {
		paymentInfo.PurchaseTime = time.Unix(0, purchase.StartTimeMillis*int64(time.Millisecond))
	}

	if purchase.UserCancellationTimeMillis > 0 {
		cancelTime := time.Unix(0, purchase.UserCancellationTimeMillis*int64(time.Millisecond))
		paymentInfo.CancellationTime = &cancelTime
	}

	if purchase.PaymentState != nil {
		paymentInfo.PaymentState = *purchase.PaymentState

		switch *purchase.PaymentState {
		case 0:
			paymentInfo.Status = receiptDomain.PaymentStatusPending
			paymentInfo.Message = "Payment is pending"
		case 1:
			// Handle cancel reason
			if purchase.CancelReason != 0 {
				switch purchase.CancelReason {
				case 0:
					paymentInfo.Status = receiptDomain.PaymentStatusCancelled
					paymentInfo.Message = "Subscription cancelled by user"
				case 1:
					paymentInfo.Status = receiptDomain.PaymentStatusCancelled
					paymentInfo.Message = "Subscription cancelled by system"
				case 2:
					paymentInfo.Status = receiptDomain.PaymentStatusRefunded
					paymentInfo.Message = "Subscription replaced with new subscription"
				case 3:
					paymentInfo.Status = receiptDomain.PaymentStatusCancelled
					paymentInfo.Message = "Subscription cancelled by developer"
				default:
					paymentInfo.Status = receiptDomain.PaymentStatusUnknown
					paymentInfo.Message = "Unknown cancel reason"
				}
			} else {
				paymentInfo.Status = receiptDomain.PaymentStatusCompleted
				paymentInfo.Message = "Payment completed and subscription is active"
			}
		case 2:
			paymentInfo.Status = receiptDomain.PaymentStatusCompleted
			paymentInfo.Message = "Free trial period"
		case 3:
			paymentInfo.Status = receiptDomain.PaymentStatusPending
			paymentInfo.Message = "Pending deferred upgrade/downgrade"
		default:
			paymentInfo.Status = receiptDomain.PaymentStatusUnknown
			paymentInfo.Message = "Unknown payment state"
		}
	} else {
		paymentInfo.Status = receiptDomain.PaymentStatusUnknown
		paymentInfo.Message = "Payment state not available"
	}

	return paymentInfo, nil
}

// IsReceiptRenewed checks all receipts and prints whether each is renewed.
func IsReceiptRenewed() {
	receipts, err := receiptDb.GetAllReceipts()
	if err != nil {
		return
	}
	for _, receipt := range receipts {
		if receipt.IsCanceled {
			continue
		}
		if time.Now().UTC().After(receipt.ExpiredAt.UTC()) {
			valid := _validateGoogleReceipt(receipt.PackageName, receipt.ProductID, receipt.PurchaseToken)

			if valid {
				stakedAmount := getTransactionAmount(receipt.ProductID)
				_, err = distributeToLineage(stakedAmount, receipt.Wallet, receipt.PurchaseToken, "")
				if err == nil {
					receiptDb.UpdateReceipt(receipt.PurchaseToken, false)
				}
			}
			continue
		}

	}
}

func _validateGoogleReceipt(packageName, productId, token string) bool {
	ctx := context.Background()

	srv, err := androidpublisher.NewService(ctx, option.WithCredentialsFile("service-account.json"))
	if err != nil {
		return false
	}

	purchase, err := srv.Purchases.Subscriptions.Get(packageName, productId, token).Do()
	if err != nil {
		return false
	}
	// Check if subscription is cancelled
	if purchase.CancelReason != 0 {
		receiptDb.UpdateReceipt(token, true)
		return false
	}

	// Check expiry time
	currentTimeMillis := time.Now().UnixNano() / int64(time.Millisecond)
	if purchase.ExpiryTimeMillis <= currentTimeMillis {
		return false
	}

	// Check payment state
	if purchase.PaymentState == nil {
		return false
	}

	// Valid payment states: 1 (received), 2 (free trial), 3 (pending deferred upgrade/downgrade)
	switch *purchase.PaymentState {
	case 0: // Payment pending
		return false
	case 1, 2, 3: // Valid states
		return true
	default: // Unknown state
		return false
	}
}
