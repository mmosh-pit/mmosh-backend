package auth

import (
	"fmt"
	"log"
	"os"

	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func ForgotPasswordVerification(email string) error {

	user, err := authDb.GetUserByEmail(email)

	if err != nil || user.ID == nil {
		return auth.ErrUserNotExists
	}

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	code := generateCode()

	existingCode := authDb.GetTemporalCode(code)

	log.Printf("Existing code: %v\n", existingCode)
	if existingCode == nil {
		from := mail.NewEmail("Kinship Bots", "security@kinshipbots.com")
		subject := "Verification Code"
		to := mail.NewEmail("", email)
		htmlContent := fmt.Sprintf("Hey there!<br /> Here's your code to reset your Password<br /> <strong>%d</strong>", code)
		message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
		res, err := client.Send(message)
		if err != nil {
			log.Println(err)
			return err
		}

		log.Printf("Response for email: %v\n", res.Body)

		authDb.SaveTemporalCode(email, code)

		return nil
	}

	return RequestCode(email)
}
