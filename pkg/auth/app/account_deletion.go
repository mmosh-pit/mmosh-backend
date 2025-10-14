package auth

import (
	"errors"
	"fmt"
	"log"
	"os"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func AccountDeletion(params *authDomain.AccountDeletionRequest) error {

	res := auth.AddAccountDeletionRequest(params)

	if res == -1 {
		return errors.New("something-went-wrong")
	}

	if res == 0 {
		client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

		from := mail.NewEmail("Kinship Bots", "security@kinshipbots.com")
		subject := "Account Deletion Request"
		to := mail.NewEmail("", "elias.ramirez@kinship.systems")
		htmlContent := fmt.Sprintf("Hey there!<br /> Someone requested to delete their account, please checkout database to see the following record<br /> <strong>%s</strong> <br/ > The reason is: %s", params.Email, params.Reason)
		to2 := mail.NewEmail("", "david.levine@kinship.systems")
		message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
		secondMessage := mail.NewSingleEmail(from, subject, to2, "", htmlContent)
		_, err := client.Send(message)
		_, err = client.Send(secondMessage)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}
