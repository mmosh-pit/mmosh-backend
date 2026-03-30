package posts

import (
	"strings"
	"time"
)

type Post struct {
	ID        string    `json:"id,omitempty"`
	Header    string    `json:"header"`
	SubHeader string    `json:"subHeader,omitempty"`
	Tags      []string  `json:"tags,omitempty"`
	Authors   []string  `json:"authors"`
	Body      string    `json:"body"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func (p *Post) Validate() bool {
	headerValid := strings.TrimSpace(p.Header) != ""
	bodyValid := strings.TrimSpace(p.Body) != ""
	slugValid := strings.TrimSpace(p.Slug) != ""
	authorsValid := len(p.Authors) > 0

	if authorsValid {
		for _, author := range p.Authors {
			if strings.TrimSpace(author) == "" {
				authorsValid = false
				break
			}
		}
	}

	return headerValid && bodyValid && slugValid && authorsValid
}
