package wallet

import (
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	wallet "github.com/mmosh-pit/mmosh_backend/pkg/wallet/db"
)

func GetAllCoinAddressHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	response, err := wallet.GetAllCoinAddresses()

	if err != nil {
		common.SendErrorResponse(w, http.StatusInternalServerError, "")
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, response)
}
