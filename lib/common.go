package lib

import (
	"context"
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api/v2"
	_ "github.com/joho/godotenv/autoload"
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

func WriteError(err error) {
	fmt.Fprint(os.Stderr, err)
}
