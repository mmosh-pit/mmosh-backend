package receiptHttp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	receiptApp "github.com/mmosh-pit/mmosh_backend/pkg/receipt/app"
	receiptDomain "github.com/mmosh-pit/mmosh_backend/pkg/receipt/domain"
)

func SaveReceiptHandler(w http.ResponseWriter, r *http.Request) {
	// Read body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	// Parse JSON
	var data receiptDomain.SaveReceiptParams
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("error decoding payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	// Validate required fields
	if strings.TrimSpace(data.Wallet) == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "wallet is required")
		return
	}
	if strings.TrimSpace(data.PackageName) == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "package_name is required")
		return
	}
	if strings.TrimSpace(data.ProductID) == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "product_id is required")
		return
	}
	if strings.TrimSpace(data.PurchaseToken) == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "purchase_token is required")
		return
	}
	if strings.TrimSpace(data.Platform) == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "platform is required")
		return
	}

	// Platform-specific validation
	platform := strings.ToLower(data.Platform)
	if platform != "android" && platform != "ios" {
		common.SendErrorResponse(w, http.StatusBadRequest, "platform must be either 'android' or 'ios'")
		return
	}

	authToken := r.Header.Get("Authorization")
	// Save receipt (includes receipt validation inside app layer)
	receipt, err := receiptApp.SaveReceipt(&data, authToken)
	if err != nil {
		log.Printf("error saving receipt: %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// Respond success
	common.SendSuccessResponse(w, http.StatusOK, receipt)
}
