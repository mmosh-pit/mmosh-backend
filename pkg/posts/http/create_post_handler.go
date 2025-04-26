package posts

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	postsApp "github.com/mmosh-pit/mmosh_backend/pkg/posts/app"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("userId")

	if userId == "" {
		common.SendErrorResponse(w, http.StatusUnauthorized, nil)
		return
	}

	var params postsApp.CreatePostParams

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(&params)
	if err != nil {
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"Invalid Payload"})
		return
	}

	if params.Prompt == "" {
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"Invalid Payload"})
		return
	}

	if dec.More() {
		common.SendErrorResponse(w, http.StatusBadRequest, []string{"Request body must only contain a single JSON object"})
		return
	}

	params.UserId = userId

	post, err := postsApp.CreatePost(&params)
	if err != nil {
		if strings.Contains(err.Error(), "required") || strings.Contains(err.Error(), "cannot be empty") {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		log.Printf("ERROR: Failed to create post: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	common.SendSuccessResponse(w, http.StatusCreated, post)
}
