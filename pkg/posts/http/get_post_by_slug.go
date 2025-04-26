package posts

import (
	"log"
	"net/http"
	"strings"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	posts "github.com/mmosh-pit/mmosh_backend/pkg/posts/app"
)

func HandlePostBySlug(w http.ResponseWriter, r *http.Request) {
	slug := strings.TrimPrefix(r.URL.Path, "/posts/slug/")
	if slug == "" {
		http.Error(w, "Missing slug in URL path after /posts/slug/", http.StatusBadRequest)
		return
	}

	post, err := posts.GetPostBySlug(slug)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, err.Error(), http.StatusNotFound) // 404 Not Found
		} else {
			log.Printf("ERROR: Failed to retrieve post by slug '%s': %v", slug, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, post)
}
