package apple

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	appleDomain "github.com/mmosh-pit/mmosh_backend/pkg/apple/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

type StoreConfig struct {
	KeyContent         []byte         // Loads a .p8 certificate
	KeyId              string         // Your private key ID from App Store Connect (Ex: 2X9R4HXF34)
	BundleId           string         // Your app’s bundle ID
	Issuer             string         // Your issuer ID from the Keys page in App Store Connect (Ex: "57246542-96fe-1a63-e053-0824d011072a")
	Sandbox            bool           // default is Production
	TokenIssuedAtFunc  func() int64   // The token’s creation time func. Default is current timestamp.
	TokenExpiredAtFunc func() int64   // The token’s expiration time func. Default is one hour later.
	TrustedCertPool    *x509.CertPool // The pool of trusted root certificates. Default is a pool containing only Apple Root CA - G3.
}

type StoreClient struct {
	Token   *Token
	cert    *Cert
	hostUrl string
}

var (
	AppStoreServerClient  *StoreClient
	AppStoreConnectClient *StoreClient
)

func InitializeAppleAppStoreClients() {
	appleAppStoreBundleId, appleAppStoreIssuer, appleAppStoreSandbox := config.GetAppleAppStoreEnvVariables()
	appleAppStoreServerPrivateKey, appleAppStoreServerKeyId := config.GetAppleAppStoreServerEnvVariables()
	appleAppStoreConnectPrivateKey, appleAppStoreConnectKeyId := config.GetAppleAppStoreConnectEnvVariables()

	storeServerConfig := &StoreConfig{
		KeyContent: []byte(appleAppStoreServerPrivateKey),
		KeyId:      appleAppStoreServerKeyId,
		BundleId:   appleAppStoreBundleId,
		Issuer:     appleAppStoreIssuer,
		Sandbox:    appleAppStoreSandbox,
	}
	storeServerToken := &Token{}
	storeServerToken.WithConfig(storeServerConfig)

	appStoreServerHostUrl := "https://api.storekit.itunes.apple.com"
	if appleAppStoreSandbox {
		appStoreServerHostUrl = "https://api.storekit-sandbox.itunes.apple.com"
	}

	AppStoreServerClient = &StoreClient{
		Token:   storeServerToken,
		cert:    newCert(storeServerConfig.TrustedCertPool),
		hostUrl: appStoreServerHostUrl,
	}

	storeConnectConfig := &StoreConfig{
		KeyContent: []byte(appleAppStoreConnectPrivateKey),
		KeyId:      appleAppStoreConnectKeyId,
		BundleId:   appleAppStoreBundleId,
		Issuer:     appleAppStoreIssuer,
		Sandbox:    appleAppStoreSandbox,
	}
	storeConnectToken := &Token{}
	storeConnectToken.WithConfig(storeConnectConfig)

	AppStoreConnectClient = &StoreClient{
		Token:   storeConnectToken,
		cert:    newCert(storeConnectConfig.TrustedCertPool),
		hostUrl: "https://api.appstoreconnect.apple.com",
	}
}

func (c *StoreClient) ParseSignedPayload(tokenStr string, claims jwt.Claims) error {
	_, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (any, error) {
		return c.cert.extractPublicKeyFromToken(tokenStr)
	})

	return err
}

type TransactionInfoResponse struct {
	SignedTransactionInfo string `json:"signedTransactionInfo"`
}

type TransactionInfoResponseError struct {
	ErrorCode    int    `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

// GetTransactionInfo https://developer.apple.com/documentation/appstoreserverapi/get_transaction_info
func (c *StoreClient) GetTransactionInfo(transactionId string) (*appleDomain.JWSTransaction, error) {
	authToken, err := c.Token.GenerateIfExpired(false)
	if err != nil {
		return nil, fmt.Errorf("GetTransactionInfo.GenerateIfExpired: %w", err)
	}

	requestUrl := c.hostUrl + fmt.Sprintf("/inApps/v1/transactions/%s", transactionId)
	req, err := http.NewRequest(http.MethodGet, requestUrl, nil)
	if err != nil {
		log.Println("GetTransactionInfo.NewRequest: ", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+authToken)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("GetTransactionInfo.Do: ", err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		if response.StatusCode == http.StatusTooManyRequests {
			retryAfter, err := strconv.ParseInt(response.Header.Get("Retry-After"), 10, 64)
			if err != nil {
				log.Println("GetTransactionInfo.ParseInt: ", err)
				return nil, err
			}

			log.Printf("Retry TransactionId: %s, at: %v\n", transactionId, time.UnixMilli(retryAfter))
			return nil, nil
		} else {
			body, err := io.ReadAll(response.Body)
			if err != nil {
				log.Println("GetTransactionInfo.ReadAll 1: ", err)
				return nil, err
			}

			log.Printf("Got body here: %v\n", string(body))

			var responseBody *TransactionInfoResponseError
			if err := json.Unmarshal(body, &responseBody); err != nil {
				log.Println("GetTransactionInfo.Unmarshal 1: ", err)
				return nil, err
			}

			return nil, fmt.Errorf("GetTransactionInfo.Response.Error: Code %d, %s", responseBody.ErrorCode, responseBody.ErrorMessage)
		}
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println("GetTransactionInfo.ReadAll 2: ", err)
		return nil, err
	}

	log.Printf("Transaction info details 1: %v\n", string(body))

	var responseBody *TransactionInfoResponse
	if err := json.Unmarshal(body, &responseBody); err != nil {
		log.Println("GetTransactionInfo.Unmarshal 2: ", err)
		return nil, err
	}

	var decodedTransaction appleDomain.JWSTransaction
	if err := c.ParseSignedPayload(responseBody.SignedTransactionInfo, &decodedTransaction); err != nil {
		log.Println("GetTransactionInfo.ParseSignedPayload: ", err)
		return nil, err
	}

	return &decodedTransaction, nil
}
