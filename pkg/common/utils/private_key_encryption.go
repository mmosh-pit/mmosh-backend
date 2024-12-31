package common

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha512"
	"encoding/hex"
	"log"

	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func EncryptPrivateKey(plaintext string) string {
	secretKey, secretIv := config.GetEncryptionKeys()

	// Derive key from secretKey using SHA512
	key := sha512.Sum512([]byte(secretKey))

	// Derive initialization vector (IV) from secretIv using SHA512
	iv := sha512.Sum512([]byte(secretIv))

	// Pad the plaintext
	plaintextBytes := PKCS5Padding([]byte(plaintext), aes.BlockSize)

	block, err := aes.NewCipher(key[:32])
	if err != nil {
		log.Printf("Got error: %v\n", err)
		return ""
	}

	ciphertext := make([]byte, len(plaintextBytes))
	mode := cipher.NewCBCEncrypter(block, iv[:16])
	mode.CryptBlocks(ciphertext, plaintextBytes)

	result := hex.EncodeToString(ciphertext)

	return result
}

func DecryptPrivateKey(encryptedData string) string {
	secretKey, secretIv := config.GetEncryptionKeys()

	// Derive key from secretKey using SHA512
	key := sha512.Sum512([]byte(secretKey))

	// Derive initialization vector (IV) from secretIv using SHA512
	iv := sha512.Sum512([]byte(secretIv))

	plaintextBytes, err := hex.DecodeString(encryptedData)

	if err != nil {
		return ""
	}

	// Create AES decrypter
	block, err := aes.NewCipher(key[:32])
	if err != nil {
		return ""
	}

	decrypted := make([]byte, len(plaintextBytes))
	mode := cipher.NewCBCDecrypter(block, iv[:16])
	mode.CryptBlocks(decrypted, plaintextBytes)

	result := hex.EncodeToString(decrypted)

	return result
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
