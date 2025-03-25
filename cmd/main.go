package main

import (
	"fmt"
	"log"
	"net/http"

	apple "github.com/mmosh-pit/mmosh_backend/pkg/apple/app"
	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/app"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
	subscriptions "github.com/mmosh-pit/mmosh_backend/pkg/subscriptions/app"
)

func main() {
	config.ValidateEnvironmentVariables("./.env")

	config.InitializeMongoConnection()

	defer config.DisconnectMongoClient()

	subscriptions.AddSubscriptionsIfNotCreatedAlready()

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
