package config

import (
	"encoding/base64"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	mongoDbURI        string
	mongoDatabaseName string
	secretKey         string
	secretIv          string

	walletBackenUrl string

	AppleAppStoreBundleId string
	AppleAppStoreIssuer   string
	AppleAppStoreSandbox  bool
	AppleAppStoreAppId    string

	AppleAppStoreServerPrivateKey string
	AppleAppStoreServerKeyId      string

	AppleAppStoreConnectPrivateKey string
	AppleAppStoreConnectKeyId      string

	GoogleBillingPubSubSubscription      string
	GoogleBillingPubSubVerificationToken string

	GoogleAppStoreBundleId string
)

func ValidateEnvironmentVariables(path string) {
	err := godotenv.Load(path)

	if err != nil {
		log.Fatalf("Could not load env variables: %v\n", err)
		return
	}

	foundMongoURI, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		panic("Missing MONGO_URI env variable")
	}

	foundDatabaseName, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		panic("Missing DATABASE_NAME env variable")
	}

	foundSecretKey, ok := os.LookupEnv("SECRET_KEY")
	if !ok {
		panic("Missing SECRET_KEY env variable")
	}

	foundSecretIv, ok := os.LookupEnv("SECRET_IV")
	if !ok {
		panic("Missing SECRET_IV env variable")
	}

	foundAppleAppStoreBundleId, ok := os.LookupEnv("APPLE_APP_STORE_BUNDLE_ID")
	if !ok {
		panic("APPLE_APP_STORE_BUNDLE_ID is not present")
	}
	foundAppleAppStoreIssuer, ok := os.LookupEnv("APPLE_APP_STORE_ISSUER")
	if !ok {
		panic("APPLE_APP_STORE_ISSUER is not present")
	}
	foundAppleAppStoreSandbox, ok := os.LookupEnv("APPLE_APP_STORE_SANDBOX")
	if !ok {
		panic("APPLE_APP_STORE_SANDBOX is not present")
	}
	boolAppleAppStoreSandbox, err := strconv.ParseBool(foundAppleAppStoreSandbox)
	if err != nil {
		panic("Error during conversion of APPLE_APP_STORE_SANDBOX env")
	}
	foundAppleAppStoreAppId, ok := os.LookupEnv("APPLE_APP_STORE_APP_ID")
	if !ok {
		panic("APPLE_APP_STORE_APP_ID is not present")
	}

	foundAppleAppStoreServerPrivateKey, ok := os.LookupEnv("APPLE_APP_STORE_SERVER_PRIVATE_KEY")
	if !ok {
		panic("APPLE_APP_STORE_SERVER_PRIVATE_KEY is not present")
	}
	foundAppleAppStoreServerKeyId, ok := os.LookupEnv("APPLE_APP_STORE_SERVER_KEY_ID")
	if !ok {
		panic("APPLE_APP_STORE_SERVER_KEY_ID is not present")
	}

	foundAppleAppStoreConnectPrivateKey, ok := os.LookupEnv("APPLE_APP_STORE_CONNECT_PRIVATE_KEY")
	if !ok {
		panic("APPLE_APP_STORE_CONNECT_PRIVATE_KEY is not present")
	}
	foundAppleAppStoreConnectKeyId, ok := os.LookupEnv("APPLE_APP_STORE_CONNECT_KEY_ID")
	if !ok {
		panic("APPLE_APP_STORE_CONNECT_KEY_ID is not present")
	}

	mongoDbURI = foundMongoURI
	mongoDatabaseName = foundDatabaseName
	secretKey = foundSecretKey
	secretIv = foundSecretIv

	decodedAppleAppStoreServerPrivateKey, err := base64.StdEncoding.DecodeString(foundAppleAppStoreServerPrivateKey)
	if err != nil {
		panic(err)
	}

	decodedAppleAppStoreConnectPrivateKey, err := base64.StdEncoding.DecodeString(foundAppleAppStoreConnectPrivateKey)
	if err != nil {
		panic(err)
	}

	foundGoogleBillingPubSubSubscription, ok := os.LookupEnv("GOOGLE_BILLING_PUBSUB_SUBSCRIPTION")
	if !ok {
		panic("GOOGLE_BILLING_PUBSUB_SUBSCRIPTION is not present")
	}
	foundGoogleBillingPubSubVerificationToken, ok := os.LookupEnv("GOOGLE_BILLING_PUBSUB_VERIFICATION_TOKEN")
	if !ok {
		panic("GOOGLE_BILLING_PUBSUB_VERIFICATION_TOKEN is not present")
	}
	foundGoogleAppStoreBundleId, ok := os.LookupEnv("GOOGLE_APP_STORE_BUNDLE_ID")
	if !ok {
		panic("GOOGLE_APP_STORE_BUNDLE_ID is not present")
	}

	foundWalletBackendUrl, ok := os.LookupEnv("WALLET_BACKEND_URL")

	if !ok {
		panic("WALLET_BACKEND_URL is missing")
	}

	AppleAppStoreBundleId = foundAppleAppStoreBundleId
	AppleAppStoreIssuer = foundAppleAppStoreIssuer
	AppleAppStoreSandbox = boolAppleAppStoreSandbox
	AppleAppStoreAppId = foundAppleAppStoreAppId

	AppleAppStoreServerPrivateKey = string(decodedAppleAppStoreServerPrivateKey)
	AppleAppStoreServerKeyId = foundAppleAppStoreServerKeyId

	AppleAppStoreConnectPrivateKey = string(decodedAppleAppStoreConnectPrivateKey)
	AppleAppStoreConnectKeyId = foundAppleAppStoreConnectKeyId

	GoogleBillingPubSubSubscription = foundGoogleBillingPubSubSubscription
	GoogleBillingPubSubVerificationToken = foundGoogleBillingPubSubVerificationToken
	GoogleAppStoreBundleId = foundGoogleAppStoreBundleId

	walletBackenUrl = foundWalletBackendUrl
}

func GetMongoURI() *string {
	return &mongoDbURI
}

func GetDatabaseName() string {
	return mongoDatabaseName
}

func GetEncryptionKeys() (string, string) {
	return secretKey, secretIv
}

func GetAppleAppStoreEnvVariables() (string, string, bool) {
	return AppleAppStoreBundleId, AppleAppStoreIssuer, AppleAppStoreSandbox
}

func GetAppleAppStoreServerEnvVariables() (string, string) {
	return AppleAppStoreServerPrivateKey, AppleAppStoreServerKeyId
}

func GetAppleAppStoreConnectEnvVariables() (string, string) {
	return AppleAppStoreConnectPrivateKey, AppleAppStoreConnectKeyId
}

func GetWalletBackendUrl() *string {
	return &walletBackenUrl
}
