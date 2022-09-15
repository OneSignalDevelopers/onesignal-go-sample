package osHelpers

import (
	"onesignal"
	"os"
)

func pushNotification() {
	apiClient := osApiClient()
	notification := *apiClient.NewNotification(AppId)
	resp, r, err :=
		.DefaultApi
		.CreateNotification(osHelpers.osAuthCtx())
		.Notification(notification).Execute()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CreateNotification`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		return
	}

	fmt.Fprintf(os.Stdout, "Response from `CreateNotification`: %v\n", resp)
}

