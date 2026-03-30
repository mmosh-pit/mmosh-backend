package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"

	authDb "github.com/mmosh-pit/mmosh_backend/pkg/auth/db"
	authDomain "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	authUtils "github.com/mmosh-pit/mmosh_backend/pkg/auth/utils"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

var testServer *httptest.Server
var walletMockServer *httptest.Server

func TestMain(m *testing.M) {
	testDatabaseURL := os.Getenv("TEST_DATABASE_URL")
	if testDatabaseURL == "" {
		fmt.Println("Skipping e2e tests: TEST_DATABASE_URL is not set")
		os.Exit(0)
	}

	// Change to project root so migrations/ and serve() work correctly.
	if err := os.Chdir(".."); err != nil {
		fmt.Printf("Failed to chdir to project root: %v\n", err)
		os.Exit(1)
	}

	// Mock wallet backend - signup calls this to create a wallet.
	walletMockServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		walletData := `{"address":"test_wallet_address","key_package":["key1","backup_key_value"]}`
		resp := map[string]any{
			"status":  true,
			"message": "ok",
			"data":    walletData,
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}))
	defer walletMockServer.Close()

	dummyB64 := base64.StdEncoding.EncodeToString([]byte("dummy-private-key-content"))
	envContent := fmt.Sprintf(
		"DATABASE_URL=%s\n"+
			"SECRET_KEY=0123456789abcdef0123456789abcdef\n"+
			"SECRET_IV=0123456789abcdef\n"+
			"APPLE_APP_STORE_BUNDLE_ID=com.test.app\n"+
			"APPLE_APP_STORE_ISSUER=test-issuer-id\n"+
			"APPLE_APP_STORE_SANDBOX=true\n"+
			"APPLE_APP_STORE_APP_ID=123456789\n"+
			"APPLE_APP_STORE_SERVER_PRIVATE_KEY=%s\n"+
			"APPLE_APP_STORE_SERVER_KEY_ID=test-server-key-id\n"+
			"APPLE_APP_STORE_CONNECT_PRIVATE_KEY=%s\n"+
			"APPLE_APP_STORE_CONNECT_KEY_ID=test-connect-key-id\n"+
			"GOOGLE_BILLING_PUBSUB_SUBSCRIPTION=test-subscription\n"+
			"GOOGLE_BILLING_PUBSUB_VERIFICATION_TOKEN=test-pubsub-token\n"+
			"GOOGLE_APP_STORE_BUNDLE_ID=com.test.app\n"+
			"WALLET_BACKEND_URL=%s\n"+
			"KARTRA_APP_ID=test-kartra-app-id\n"+
			"KARTRA_API_KEY=test-kartra-api-key\n"+
			"KARTRA_API_PASSWORD=test-kartra-password\n"+
			"KARTRA_API_BASE=http://localhost:19999\n"+
			"AI_API_URL=http://localhost:19999\n"+
			"OPEN_AI_KEY=test-openai-key\n"+
			"NEXT_BACKEND_URL=http://localhost:19999\n"+
			"AI_API_BASE=http://localhost:19999\n"+
			"ANTHROPIC_KEY=test-anthropic-key\n"+
			"SENDGRID_API_KEY=test-sendgrid-key\n",
		testDatabaseURL, dummyB64, dummyB64, walletMockServer.URL,
	)

	tmpEnv, err := os.CreateTemp("", "e2e-*.env")
	if err != nil {
		fmt.Printf("Failed to create temp env file: %v\n", err)
		os.Exit(1)
	}
	tmpEnv.WriteString(envContent)
	tmpEnv.Close()
	defer os.Remove(tmpEnv.Name())

	config.ValidateEnvironmentVariables(tmpEnv.Name())
	config.InitializePostgresConnection()
	defer config.DisconnectPostgresClient()

	mg, err := migrate.New("file://migrations", testDatabaseURL)
	if err != nil {
		fmt.Printf("Failed to initialize migrations: %v\n", err)
		os.Exit(1)
	}
	if upErr := mg.Up(); upErr != nil && upErr != migrate.ErrNoChange {
		fmt.Printf("Failed to run migrations: %v\n", upErr)
		mg.Close()
		os.Exit(1)
	}
	mg.Close()

	testServer = httptest.NewServer(http.HandlerFunc(serve))
	defer testServer.Close()

	os.Exit(m.Run())
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

// doRequest performs an HTTP request against the test server and returns the
// decoded response body, the raw body bytes, and the HTTP status code.
func doRequest(method, path string, body any, token string) (*http.Response, []byte) {
	var reqBody io.Reader
	if body != nil {
		data, _ := json.Marshal(body)
		reqBody = bytes.NewReader(data)
	}

	req, err := http.NewRequestWithContext(context.Background(), method, testServer.URL+path, reqBody)
	if err != nil {
		panic(fmt.Sprintf("failed to build request: %v", err))
	}
	req.Header.Set("Content-Type", "application/json")
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(fmt.Sprintf("request failed: %v", err))
	}
	defer resp.Body.Close()
	raw, _ := io.ReadAll(resp.Body)
	return resp, raw
}

// decodeData unmarshals the {"data": ...} wrapper and returns the inner value.
func decodeData(raw []byte, out any) error {
	var wrapper map[string]json.RawMessage
	if err := json.Unmarshal(raw, &wrapper); err != nil {
		return err
	}
	return json.Unmarshal(wrapper["data"], out)
}

// createTestUser inserts a user directly into the DB and returns (userID, sessionToken).
// This bypasses wallet/email flows so that auth tests don't depend on external services.
func createTestUser(email, role string) (string, string) {
	pw, err := authUtils.EncryptPassword("TestPass123!")
	if err != nil {
		panic(fmt.Sprintf("encrypt password: %v", err))
	}

	token, err := authUtils.GenerateSessionToken([]string{})
	if err != nil {
		panic(fmt.Sprintf("generate token: %v", err))
	}

	user := &authDomain.User{
		Name:      "Test User",
		Email:     email,
		Password:  pw,
		Sessions:  []string{*token},
		UUID:      uuid.New().String(),
		Wallet:    "test_wallet_" + uuid.New().String(),
		Picture:   "https://storage.googleapis.com/mmosh-assets/default.png",
		Role:      role,
		FromBot:   "KIN",
		CreatedAt: time.Now(),
	}

	if err := authDb.CreateUser(user); err != nil {
		panic(fmt.Sprintf("create test user (%s): %v", email, err))
	}

	return user.ID, *token
}

// uniqueEmail returns a unique email address for each test run.
func uniqueEmail(prefix string) string {
	return fmt.Sprintf("%s-%d@example.com", prefix, time.Now().UnixNano())
}
