package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

type ResponseModel struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func CreateWallet(email string) (string, error) {
	wallet := authDb.GetWalletByEmail(email)

	if wallet != nil {
		log.Println("Already got wallet!!!")
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

	authDb.SaveWalletToDb(email, &walletData)

	return walletData.Address, nil
}
