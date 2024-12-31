package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	mongoDbURI        string
	mongoDatabaseName string
	secretKey         string
	secretIv          string
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

	mongoDbURI = foundMongoURI
	mongoDatabaseName = foundDatabaseName
	secretKey = foundSecretKey
	secretIv = foundSecretIv
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
