package google

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	googleApp "github.com/mmosh-pit/mmosh_backend/pkg/google/app"
	googleDomain "github.com/mmosh-pit/mmosh_backend/pkg/google/domain"
)

func decodeAndValidate(msg googleDomain.PushSubscription) (*googleDomain.MessageData, string, error) {
	if msg.Subscription != config.GoogleBillingPubSubSubscription {
		return nil, "", errors.New("bad subscription")
	}

	decodedData, err := base64.StdEncoding.DecodeString(msg.Message.Data)
	if err != nil {
		return nil, "", fmt.Errorf("error decoding base64 data: %w", err)
	}

	var messageData googleDomain.MessageData
	err = json.Unmarshal(decodedData, &messageData)
	if err != nil {
		return nil, "", fmt.Errorf("error unmarshalling JSON: %w", err)
	}

	var notificationType string
	switch {
	case messageData.SubscriptionNotification != (googleDomain.SubscriptionNotification{}):
		notificationType = "SubscriptionNotification"
	case messageData.OneTimeProductNotification != (googleDomain.OneTimeProductNotification{}):
		notificationType = "OneTimeProductNotification"
	case messageData.VoidedPurchaseNotification != (googleDomain.VoidedPurchaseNotification{}):
		notificationType = "VoidedPurchaseNotification"
	case messageData.TestNotification != (googleDomain.TestNotification{}):
		notificationType = "TestNotification"
	default:
		return nil, "", fmt.Errorf("unknown notification type")
	}

	return &messageData, notificationType, nil
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("token") != config.GoogleBillingPubSubVerificationToken {
		log.Println("Google bad token")
		http.Error(w, "Bad token", http.StatusBadRequest)
		return
	}

	msg := &googleDomain.PushSubscription{}
	if err := json.NewDecoder(r.Body).Decode(msg); err != nil {
		log.Println("Could not decode google body: ", err)
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	messageData, notificationType, err := decodeAndValidate(*msg)
	if err != nil {
		log.Printf("Error in Google decodeAndValidate with msg: %v, %v\n", msg, err)
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("messageData: ", *messageData)

	switch notificationType {
	case "SubscriptionNotification":
		err := googleApp.HandleSubscriptionNotification(messageData.PackageName, messageData.SubscriptionNotification)
		if err != nil {
			log.Printf("Error in HandleSubscriptionNotification with: %v, %s\n", messageData.SubscriptionNotification, err.Error())
		}
	case "TestNotification":
		log.Println("Received a TestNotification from Google RTDN!")
	}
}
