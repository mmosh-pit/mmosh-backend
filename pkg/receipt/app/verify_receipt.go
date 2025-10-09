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

	fmt.Println("----- Payment state -----:", purchase.PaymentState)
	if purchase.PaymentState != nil {
		paymentInfo.PaymentState = *purchase.PaymentState
		fmt.Println("----- Payment Info -----:", paymentInfo)

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
