package notificationApp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	notificationDomain "github.com/mmosh-pit/mmosh_backend/pkg/notification/domain"
)

func PushOneSignalNotification(title string, message string, playerIds []string) (*notificationDomain.OneSignalResponse, error) {
	if title == "" || message == "" {
		return nil, fmt.Errorf("title and message cannot be empty")
	}
	if len(playerIds) == 0 {
		return nil, fmt.Errorf("playerIds cannot be empty")
	}

	apiKey := os.Getenv("ONESIGNAL_API_KEY")
	appID := os.Getenv("ONESIGNAL_APP_ID")

	if apiKey == "" {
		return nil, fmt.Errorf("ONESIGNAL_API_KEY environment variable not set")
	}
	if appID == "" {
		return nil, fmt.Errorf("ONESIGNAL_APP_ID environment variable not set")
	}

	if len(appID) > 8 {
		log.Printf("Using OneSignal App ID: %s...", appID[:8])
	} else {
		log.Printf("Using OneSignal App ID: %s", appID)
	}

	body := map[string]interface{}{
		"app_id":             appID,
		"include_player_ids": playerIds,
		"headings":           map[string]string{"en": title},
		"contents":           map[string]string{"en": message},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Basic %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		var errorResp map[string]interface{}
		if err := json.Unmarshal(respBody, &errorResp); err == nil {
			if errors, ok := errorResp["errors"]; ok {
				return nil, fmt.Errorf("OneSignal API error: %v", errors)
			}
		}
		return nil, fmt.Errorf("OneSignal API error (status %d): %s", resp.StatusCode, string(respBody))
	}

	var oneSignalResp notificationDomain.OneSignalResponse
	if err := json.Unmarshal(respBody, &oneSignalResp); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}
	return &oneSignalResp, nil
}
