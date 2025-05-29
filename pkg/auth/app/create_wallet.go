package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type ResponseModel struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func CreateWallet(email string) (string, error) {
	mailClient := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	wallet := authDb.GetWalletByEmail(email)

	if wallet != nil {
		return wallet.Address, authDomain.ErrWalletAlreadyExists
	}

	client := &http.Client{}

	baseUrl := config.GetWalletBackendUrl()

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/create", *baseUrl), nil)

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Got error trying to create wallet on request: %v\n", err)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Got error reading wallet response: %v\n", err)
		return "", err
	}

	var response ResponseModel

	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Printf("Got error decoding wallet response: %v\n", err)
		return "", err
	}

	if !response.Status {
		return "", authDomain.ErrSomethingWentWrong
	}

	var walletData authDomain.WalletResponse

	err = json.Unmarshal([]byte(response.Data), &walletData)

	if err != nil {
		log.Printf("Got error on final decoding for wallet: %v\n", err)
		return "", err
	}

	mailKeyPair := walletData.KeyPackage[1]

	from := mail.NewEmail("Kinship Bots", "support@liquidhearts.club")
	subject := "Kinship Wallet Key Pair"
	to := mail.NewEmail("", email)
	htmlContent := fmt.Sprintf("Here's your Kinship wallet keypair, save it in a safe place in case you need to recover your Kinship Wallet<br /><br /> <strong>%s</strong>", mailKeyPair)
	message := mail.NewSingleEmail(from, subject, to, "", htmlContent)
	res, err := mailClient.Send(message)

	if err != nil {
		log.Printf("Error trying to send keypair to email: %v\n", err)
		log.Printf("Wallet: %v\n", mailKeyPair)
		authDb.SaveFailedEmailAttemptsKeypairs(email, mailKeyPair)
	}

	log.Printf("Send wallet keypair result: %v\n", res)

	authDb.SaveWalletToDb(email, &walletData)

	return walletData.Address, nil
}
