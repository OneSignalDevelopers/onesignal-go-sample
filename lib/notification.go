package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api/v2"
)

var configuration = onesignal.NewConfiguration()
var apiClient = onesignal.NewAPIClient(configuration)

func PushNotification() {
	appId := os.Getenv("APP_ID")
	restApiKey := os.Getenv("REST_API_KEY")
	osAuthCtx := context.WithValue(
		context.Background(),
		onesignal.AppAuth,
		restApiKey,
	)

	notification := *onesignal.NewNotification(appId)
	notification.SetIncludedSegments([]string{"Subscribed Users"})
	notification.SetIsIos(false)
	message := "Go Test Notification"
	stringMap := onesignal.StringMap{En: &message}
	notification.Contents = *onesignal.NewNullableStringMap(&stringMap)

	resp, r, err := apiClient.DefaultApi.CreateNotification(osAuthCtx).Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateNotification`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `CreateNotification`: %v\n", resp)
	fmt.Fprintf(os.Stdout, "Notification ID: %v\n", resp.GetId())
}
