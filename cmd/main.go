package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/robfig/cron/v3"

	apple "github.com/mmosh-pit/mmosh_backend/pkg/apple/app"
	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/app"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	receiptApp "github.com/mmosh-pit/mmosh_backend/pkg/receipt/app"
	subscriptions "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/app"
)

func main() {
	config.ValidateEnvironmentVariables("./.env")

	config.InitializeMongoConnection()

	defer config.DisconnectMongoClient()

	subscriptions.AddSubscriptionsIfNotCreatedAlready()

	c := cron.New()
	c.AddFunc("@every 1h", func() {
		log.Println("Cron job running every 1 hour")
		receiptApp.IsReceiptRenewed()
	})
	c.Start()
	defer c.Stop()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 6050),
		Handler: http.HandlerFunc(serve),
	}

	log.Printf("Server starting on port 6050")

	go chat.CreatePool()
	go chat.SetupBotUsers()
	go apple.InitializeAppleAppStoreClients()

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}
