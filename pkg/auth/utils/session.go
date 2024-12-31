package auth

import (
	"crypto/rand"
	"math/big"
)

func GenerateSessionToken(sessions []string) (*string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, 32)
	for i := 0; i < 32; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return nil, err
		}
		ret[i] = letters[num.Int64()]
	}
	token := string(ret)

	exists := false

	for _, val := range sessions {
		if token == val {
			exists = true
		}
	}

	if exists {
		return GenerateSessionToken(sessions)
	}

	return &token, nil
}
