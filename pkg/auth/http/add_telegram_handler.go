package auth

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddTelegramHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("error reading payload: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	var data authDomain.TelegramUserData
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Printf("error decoding payload on bluesky: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	userIdBson, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Printf("error transforming to Object ID telegram: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	err = auth.SaveTelegramData(data, userIdBson)

	if err != nil {
		switch {
		case errors.Is(err, authDomain.ErrInvalidBluesky):
			common.SendErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		default:
			common.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		}
	}

	common.SendSuccessResponse(w, http.StatusCreated, nil)
}
