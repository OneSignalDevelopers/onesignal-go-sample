package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api"
)

var configuration = onesignal.NewConfiguration()
var apiClient = onesignal.NewAPIClient(configuration)
var osAuthCtx = context.WithValue(
	context.Background(),
	onesignal.AppAuth,
	os.Getenv("REST_API_KEY"),
)

func PushNotification() {
	fmt.Fprintf(os.Stdout, "oAuthCtx: %v\n", osAuthCtx)

	notification := *onesignal.NewNotification(os.Getenv("APP_ID"))
	notification.SetIncludedSegments([]string{"Subscribed Users"})

	resp, r, err := apiClient.DefaultApi.CreateNotification(osAuthCtx).Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateNotification`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `CreateNotification`: %v\n", resp)
}
