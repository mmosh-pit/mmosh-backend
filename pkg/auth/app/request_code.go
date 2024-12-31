package auth

import (
	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
)

type RequestCodeParams struct {
	Email string `json:"email"`
}

func RequestCode(email string) error {

	// client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	code := generateCode()

	existingCode := authDb.GetTemporalCode(code)

	if existingCode == nil {
		// from := mail.NewEmail("Kinship Codes", "support@liquidhearts.club")
		// subject := "Verification Code"
		// to := mail.NewEmail("", email)
		// htmlContent := fmt.Sprintf("Hey there!<br /> Here's your code to verify your Email and finish your registration into Liquid Hearts Club!<br /> <strong>%d</strong>", code)
		// message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
		// response, err := client.Send(message)
		// if err != nil {
		// 	log.Println(err)
		// 	return err
		// }

		authDb.SaveTemporalCode(email, code)

		return nil
	}

	return RequestCode(email)
}

func generateCode() int {
	return 000000
	// return int(math.Floor(100000 + rand.Float64()*900000))
}
