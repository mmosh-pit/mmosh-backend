package notificationHttp

import (
	"encoding/json"
	"io"
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	notificationApp "github.com/mmosh-pit/mmosh_backend/pkg/notification/app"
	notificationDb "github.com/mmosh-pit/mmosh_backend/pkg/notification/db"
	notificationDomain "github.com/mmosh-pit/mmosh_backend/pkg/notification/domain"
)

func PushNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	// Parse JSON request body
	var data notificationDomain.PushNotificationRequestParams
	if err := json.Unmarshal(body, &data); err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	// Validate required fields
	if data.Title == "" || data.Message == "" || data.Wallet == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "title, message, and wallet are required")
		return
	}

	// Get player IDs from database
	playerIds, err := notificationDb.GetPlayerIds(data.Wallet)
	if err != nil {
		common.SendErrorResponse(w, http.StatusNotFound, err.Error())
		return
	}

	// Send notification via OneSignal
	oneSignalResponse, err := notificationApp.PushOneSignalNotification(data.Title, data.Message, playerIds)
	if err != nil {
		common.SendErrorResponse(w, http.StatusInternalServerError, "failed to send notification")
		return
	}

	// Send success response
	response := map[string]interface{}{
		"success": true,
		"data":    oneSignalResponse,
	}
	common.SendSuccessResponse(w, http.StatusOK, response)
}
