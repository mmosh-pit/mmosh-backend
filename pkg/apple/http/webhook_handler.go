package apple

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	appleApp "github.com/mmosh-pit/mmosh_backend/pkg/apple/app"
	appleDomain "github.com/mmosh-pit/mmosh_backend/pkg/apple/domain"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Got apple callback!!!")

	body := &appleDomain.ResponseBodyV2{}
	if err := json.NewDecoder(r.Body).Decode(body); err != nil {
		log.Println("Could not decode apple body: ", err)
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	var decodedPayload appleDomain.ResponseBodyV2DecodedPayload
	if err := appleApp.AppStoreServerClient.ParseSignedPayload(body.SignedPayload, &decodedPayload); err != nil {
		log.Println("Got error parsing payload: ", err)
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	var decodedTransaction appleDomain.JWSTransaction
	if err := appleApp.AppStoreServerClient.ParseSignedPayload(decodedPayload.Data.SignedTransactionInfo, &decodedTransaction); err != nil {
		log.Println("Got error parsing transaction: ", err)
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	transaction, err := appleApp.AppStoreServerClient.GetTransactionInfo(decodedTransaction.TransactionId)

	log.Printf("Got transaction info here: %v\n", *transaction)

	if err != nil {
		log.Println("GetTransactionInfo: ", err)
		http.Error(w, fmt.Sprintf("Invalid transaction: %v", err), http.StatusBadRequest)
		return
	}

	switch transaction.Type {
	case appleDomain.AUTO_RENEWABLE_SUBSCRIPTION, appleDomain.NON_RENEWING_SUBSCRIPTION:
		if err := appleApp.HandleSubscriptionPurchases(decodedPayload.NotificationType, decodedPayload.SubType, *transaction); err != nil {
			log.Printf("Error in HandleSubscriptionPurchases with: %v, %s\n", *transaction, err.Error())
		}
	}
}
