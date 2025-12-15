package http

import (
	"log"
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	stripeApp "github.com/mmosh-pit/mmosh_backend/pkg/stripe/app"
)

func GetStripeOnboardURLHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")
	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	url, err := stripeApp.GetOnboardURL(userId)
	if err != nil {
		log.Println("Error in GetStripeOnboardURL: ", err)
		common.SendErrorResponse(w, http.StatusInternalServerError, "something went wrong")
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, map[string]interface{}{"onboard_url": url})
}
