package mail

import (
	"log"
	"net/http"
	"strings"

	"github.com/cloudmailin/cloudmailin-go"
	mailApp "github.com/mmosh-pit/mmosh_backend/pkg/mail/app"
)

func IncomingEmailHandler(w http.ResponseWriter, r *http.Request) {
	message, err := cloudmailin.ParseIncoming(r.Body)
	if err != nil {
		http.Error(w, "Error parsing message: "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	if strings.HasPrefix(message.Envelope.To, "noreply@") {
		http.Error(w, "No replies please", http.StatusForbidden)
		return
	}

	body := message.ReplyPlain
	if body == "" {
		body = message.Plain
	}

	log.Printf("Plain: %v\n", message.Plain)
	log.Printf("From: %v\n", message.Envelope.From)

	mailApp.ProcessIncomingEmail(message.Envelope.From, message.Plain, message.Envelope.To)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
