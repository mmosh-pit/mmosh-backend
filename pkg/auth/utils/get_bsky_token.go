package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	API_URL = "https://bsky.social/xrpc"
)

type DIDDoc struct {
	Context            []string `json:"@context"`
	ID                 string   `json:"id"`
	AlsoKnownAs        []string `json:"alsoKnownAs"`
	VerificationMethod []struct {
		ID                 string `json:"id"`
		Type               string `json:"type"`
		Controller         string `json:"controller"`
		PublicKeyMultibase string `json:"publicKeyMultibase"`
	} `json:"verificationMethod"`
	Service []struct {
		ID              string `json:"id"`
		Type            string `json:"type"`
		ServiceEndpoint string `json:"serviceEndpoint"`
	} `json:"service"`
}

type DIDResponse struct {
	DID             string `json:"did"`
	DIDDoc          DIDDoc `json:"didDoc"`
	Handle          string `json:"handle"`
	Email           string `json:"email"`
	EmailConfirmed  bool   `json:"emailConfirmed"`
	EmailAuthFactor bool   `json:"emailAuthFactor"`
	AccessJwt       string `json:"accessJwt"`
	RefreshJwt      string `json:"refreshJwt"`
	Active          bool   `json:"active"`
}

func IsValidConnection(identifier, password string) bool {

	if identifier == "" {
		return false
	}

	if password == "" {
		return false
	}

	requestBody, err := json.Marshal(map[string]string{
		"identifier": identifier,
		"password":   password,
	})
	if err != nil {
		return false
	}

	url := fmt.Sprintf("%s/com.atproto.server.createSession", API_URL)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false
	}

	var tokenResponse DIDResponse
	if err := json.NewDecoder(resp.Body).Decode(&tokenResponse); err != nil {
		return false
	}

	return true
}
