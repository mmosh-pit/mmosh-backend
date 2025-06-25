package auth

import (
	"log"
	"net/http"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteBlueskyHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, "")
		return
	}

	userIdBson, err := primitive.ObjectIDFromHex(userId)

	if err != nil {
		log.Printf("error transforming to Object ID bluesky: %v", err)
		common.SendErrorResponse(w, http.StatusBadRequest, "invalid payload")
		return
	}

	auth.DeleteBlueskyData(userIdBson)

	common.SendSuccessResponse(w, http.StatusOK, nil)
}
