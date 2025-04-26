package posts

import (
	"log"
	"net/http"

	common "github.com/mmosh-pit/mmosh_backend/pkg/common/utils"
	posts "github.com/mmosh-pit/mmosh_backend/pkg/posts/app"
)

func HandlePostsByAuthor(w http.ResponseWriter, r *http.Request) {

	authors := r.URL.Query()["authors"]

	if len(authors) == 0 {
		http.Error(w, "Query parameter 'authors' is required", http.StatusBadRequest)
		return
	}

	posts, err := posts.GetPostsByAuthors(authors)
	if err != nil {
		log.Printf("ERROR: Failed to retrieve posts by authors %v: %v", authors, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	common.SendSuccessResponse(w, http.StatusOK, posts)
}
