package subscriptions

import (
	"log"
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	subscriptions "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/db"
)

func GetSubscriptionsHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	response, err := subscriptions.GetSubscriptions()

	if err != nil {
		log.Printf("Error while retrieving subscriptions: %v\n", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, "Something went wrong")
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, response)
}
