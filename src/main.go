package main

import (
	"context"
	"fmt"
	"log"
	"onesignal"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appId := os.Getenv("ONESIGNAL_APP_ID")

	notification := *openapiclient.NewNotification("AppId_example")
	configuration := onesignal.NewConfiguration()
	apiClient := onesignal.NewAPIClient(configuration)
	appAuth := context.WithValue(context.Background(), onesignal.AppAuth, appId)
	resp, r, err := apiClient.DefaultApi.CreateNotification(appAuth).Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateNotification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotification`: InlineResponse200
	fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateNotification`: %v\n", resp)
}
