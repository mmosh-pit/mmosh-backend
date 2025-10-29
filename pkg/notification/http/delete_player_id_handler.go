package notificationHttp

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	notificationDb "github.com/mmosh-pit/mmosh_backend/pkg/notification/db"
	notificationDomain "github.com/mmosh-pit/mmosh_backend/pkg/notification/domain"
)

func DeletePlayerIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete && r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get userId from header
	userId := r.Header.Get("userId")
	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "userId is required")
		return
	}

	// Read request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	// Parse JSON request body
	var data notificationDomain.InsertPlayIdRequestParams
	if err := json.Unmarshal(body, &data); err != nil {
		log.Printf("error decoding payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid JSON")
		return
	}

	// Validate playerId
	if data.PlayerId == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, "playerId is required")
		return
	}

	// Delete player ID
	success, err := notificationDb.DeletePlayerId(data.PlayerId, data.Wallet)
	if err != nil {
		log.Printf("error deleting player ID: %v", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	if !success {
		common.SendErrorResponse(w, http.StatusInternalServerError, "failed to delete player ID")
		return
	}

	// Send success response
	response := map[string]interface{}{
		"success": true,
		"message": "Player ID deleted successfully",
	}
	common.SendSuccessResponse(w, http.StatusOK, response)
}
