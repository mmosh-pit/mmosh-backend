package receiptApp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	receiptDb "github.com/mmosh-pit/mmosh_backend/pkg/receipt/db"
	receiptDomain "github.com/mmosh-pit/mmosh_backend/pkg/receipt/domain"
	"google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
)

// SaveReceipt validates the subscription and saves it in DB
func SaveReceipt(params *receiptDomain.SaveReceiptParams, authToken string) (*receiptDomain.Receipt, error) {
	receipt := &receiptDomain.Receipt{
		PackageName:   params.PackageName,
		ProductID:     params.ProductID,
		PurchaseToken: params.PurchaseToken,
		Wallet:        params.Wallet,
		Platform:      params.Platform,
		CreatedAt:     time.Now(),
	}

	valid, err := validateGoogleReceipt(params.PackageName, params.ProductID, params.PurchaseToken)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, fmt.Errorf("subscription invalid or expired")
	}

	if err := receiptDb.SaveReceipt(receipt); err != nil {
		return nil, err
	}

	// Distribute to lineage after successful receipt save
	stakedAmount := getTransactionAmount(params.ProductID)
	_, err = distributeToLineage(stakedAmount, params.Wallet, params.PurchaseToken, authToken)
	if err != nil {
		// Log the error but don't fail the entire operation
		fmt.Printf("warning: failed to distribute to lineage: %v\n", err)
		// You might want to implement retry logic here
	}

	return receipt, nil
}

// validateGoogleReceipt checks Google Play subscription status
func validateGoogleReceipt(packageName, productId, token string) (bool, error) {
	ctx := context.Background()

	srv, err := androidpublisher.NewService(ctx, option.WithCredentialsFile("service-account.json"))
	if err != nil {
		return false, fmt.Errorf("error creating android publisher service: %v", err)
	}

	purchase, err := srv.Purchases.Subscriptions.Get(packageName, productId, token).Do()
	if err != nil {
		return false, fmt.Errorf("error fetching subscription: %v", err)
	}

	fmt.Println("Receipt purchase details:", purchase)
	fmt.Println("PaymentState:", purchase.PaymentState)
	fmt.Println("AutoRenewing:", purchase.AutoRenewing)
	fmt.Println("ExpiryTimeMillis:", purchase.ExpiryTimeMillis)

	// Validate subscription based on payment state and expiry
	isValid := purchase.PaymentState != nil && *purchase.PaymentState == 1 && purchase.ExpiryTimeMillis > time.Now().UnixNano()/int64(time.Millisecond)

	if isValid {
		fmt.Println("Subscription purchase valid")
	} else {
		fmt.Println("Subscription purchase expired or invalid")
	}

	return isValid, nil
}

func getTransactionAmount(productId string) int {
	switch productId {
	case "enjoyer":
		return 15
	case "enjoyery":
		return 90
	case "creator":
		return 24
	default:
		return 180
	}
}

func distributeToLineage(stakedAmount int, userAddress string, purchaseToken string, authToken string) (*http.Response, error) {
	requestBody := receiptDomain.DistributeRequest{
		StakedAmount: stakedAmount,
		UserAddress:  userAddress,
		PurchaseID:   purchaseToken,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	backendURL := os.Getenv("NEXT_BACKEND_URL")
	if backendURL == "" {
		return nil, fmt.Errorf("NEXT_BACKEND_URL environment variable not set")
	}

	url := backendURL + "/api/distribute-to-lineage"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", authToken)

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	return client.Do(req)
}
