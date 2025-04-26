package auth

import (
	"log"
	"time"

	auth "github.com/mmosh-pit/mmosh_backend/pkg/auth/domain"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func SaveWalletToDb(email string, wallet *auth.WalletResponse) {
	log.Println("Saving wallet...")
	client, ctx := config.GetMongoClient()
	databaseName := config.GetDatabaseName()

	collection := client.Database(databaseName).Collection("mmosh-app-user-wallet")

	data := auth.Wallet{
		Address:    wallet.Address,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		KeyPackage: wallet.KeyPackage[0],
		Email:      email,
	}

	_, err := collection.InsertOne(*ctx, data)

	if err != nil {
		log.Printf("Error trying to save wallet: %v\n", err)
	}
}
