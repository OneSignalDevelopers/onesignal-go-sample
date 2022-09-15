package osHelpers

import (
	"context"
	"onesignal"
	"os"
)

func osAuthCtx() {
	context.WithValue(
		context.Background(),
		onesignal.AppAuth,
		os.Getenv("App_ID")
	)
}

func osApiClient() {
	configuration := onesignal.NewConfiguration()
	apiClient := onesignal.NewAPIClient(configuration)
}
