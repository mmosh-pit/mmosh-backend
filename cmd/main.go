package main

import (
	"fmt"
	"log"
	"net/http"

	chat "github.com/mmosh-pit/mmosh_backend/pkg/chat/app"
	"github.com/mmosh-pit/mmosh_backend/pkg/config"
)

func main() {
	config.ValidateEnvironmentVariables("./.env")

	config.InitializeMongoConnection()

	defer config.DisconnectMongoClient()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 6000),
		Handler: http.HandlerFunc(serve),
	}

	log.Printf("Server starting on port 6000")

	go chat.CreatePool()
	go chat.SetupBotUsers()

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}
