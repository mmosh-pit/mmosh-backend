package mail

import (
	"fmt"
	"log"

	"github.com/cloudmailin/cloudmailin-go"
	common "github.com/mmosh-pit/mmosh_backend/pkg/common/app"
)

func ProcessIncomingEmail(email, text, to string) {
	response := common.FetchAIResponse("VISITOR", text, "", []string{"PUBLIC"})

	if response != "" {

		// Create the default CloudMailin Client. This example will panic if there
		// are any failures at all.
		client, err := cloudmailin.NewClient()
		if err != nil {
			log.Printf("Error trying to create email client: %v\n", err)
			return
		}

		fromEmail := to

		email := cloudmailin.OutboundMail{
			From:     fromEmail,
			To:       []string{email},
			Headers:  map[string][]string{"x-agent": {"cloudmailin-go"}},
			Subject:  "AI Response",
			Plain:    response,
			HTML:     response,
			Priority: "",
			Tags:     []string{"go"},
			TestMode: true,
		}

		// This will re-write the email struct based on the
		// JSON returned from the call if successful.
		_, err = client.SendMail(&email)
		if err != nil {
			log.Printf("Got error trying to send email: %v\n", err)
			return
		}

		// The email.ID should now be populated
		fmt.Printf("ID: %s, Tags: %s", email.ID, email.Tags)

	}
}
