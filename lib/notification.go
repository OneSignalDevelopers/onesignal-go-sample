package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api/v2"
)

var config = onesignal.NewConfiguration()
var apiClient = onesignal.NewAPIClient(config)
var appId = os.Getenv("APP_ID")
var restApiKey = os.Getenv("REST_API_KEY")
var authCtx = context.WithValue(
	context.Background(),
	onesignal.AppAuth,
	restApiKey,
)

func PushNotification() {
	notification := *onesignal.NewNotification(appId)
	notification.SetIncludedSegments([]string{"Subscribed Users"})
	notification.SetIsIos(false)
	message := "Go Test Notification"
	stringMap := onesignal.StringMap{En: &message}
	notification.Contents = *onesignal.NewNullableStringMap(&stringMap)

	resp, r, err := apiClient.DefaultApi.CreateNotification(authCtx).Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateNotification`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `CreateNotification`: %v\n", resp)
	fmt.Fprintf(os.Stdout, "Notification ID: %v\n", resp.GetId())
}

// Create a user
func CreatUser() {
	u := *onesignal.NewUser()
	user, _, err := apiClient.DefaultApi.CreateUser(authCtx, appId).User(u).Execute()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	fmt.Fprintf(os.Stdout, "New user created: %v\n", user)
}

// View a user
// Update a user
// Delete a user
// View user identity
// View user identity by subscription
// Create alias
// Create alias by subscription
// Delete alias

// Create subscription
// Update subscription
// View subscription I AMs
// Deletee subscription
// Transfer subscription
