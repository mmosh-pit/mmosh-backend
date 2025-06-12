package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

type SignTransactionResponse struct {
	Signature string `json:"signature"`
	Address   string `json:"address"`
}

type SignTransactionRequest struct {
	Address string      `json:"address"`
	Message string      `json:"message"`
	Key     interface{} `json:"key"`
}

func SignTransaction(userId, message string) (*SignTransactionResponse, error) {
	user, err := authDb.GetUserById(userId)

	if err != nil {
		return nil, err
	}

	wallet := authDb.GetWalletByEmail(user.Email)

	baseUrl := config.GetWalletBackendUrl()

	reqData := SignTransactionRequest{
		Address: wallet.Address,
		Key:     wallet.Private,
		Message: message,
	}

	encoded, err := json.Marshal(reqData)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/sign", *baseUrl), bytes.NewBuffer(encoded))

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Printf("Error encoding? %v\n", err)
		return nil, err
	}

	resp, err := client.Do(req)

	if err != nil {
		log.Printf("Error doing? %v\n", err)
		return nil, err
	}

	var response ResponseModel

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Got error reading body? %s, %v\n", string(body), err)
		return nil, err
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		log.Printf("Got error unmarshalling? %s, %v\n", string(body), err)
		return nil, err
	}

	if !response.Status {
		return nil, authDomain.ErrSomethingWentWrong
	}

	var result SignTransactionResponse

	// err = json.Unmarshal([]byte(response.Data), &result)
	//
	// if err != nil {
	// 	return nil, err
	// }

	result.Signature = response.Data
	result.Address = wallet.Address

	return &result, nil
}
