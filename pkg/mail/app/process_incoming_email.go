package mail

import (
	"fmt"

	"github.com/cloudmailin/cloudmailin-go"
)

func ProcessIncomingEmail(email, text, to string) {
	response := FetchAIResponse("VISITOR", text, []string{"PUBLIC"})

	if response != "" {

		// Create the default CloudMailin Client. This example will panic if there
		// are any failures at all.
		client, err := cloudmailin.NewClient()
		if err != nil {
			panic(err)
		}

		fromEmail := to

		// SMTP Settings will be taken from CLOUDMAILIN_SMTP_URL env variable by
		// default but they can be overridden.
		// client.SMTPAccountID = ""
		// client.SMTPToken = ""

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
			panic(err)
		}

		// The email.ID should now be populated
		fmt.Printf("ID: %s, Tags: %s", email.ID, email.Tags)

	}
}
