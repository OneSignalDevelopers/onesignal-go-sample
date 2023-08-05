package lib

import (
	"fmt"
	"os"

	"github.com/OneSignal/onesignal-go-api/v2"
)

func CreatUser() {
	u := *onesignal.NewUser()
	user, res, err := apiClient.DefaultApi.CreateUser(authCtx, appId).User(u).Execute()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	fmt.Fprintf(os.Stdout, "Create User Response: %v\n", res)
	fmt.Fprintf(os.Stdout, "User: %v\n", user)
}

func ViewUser(aliasLabel string, aliasId string) {

	fmt.Println(appId)
	user, res, err := apiClient.DefaultApi.FetchUser(authCtx, appId, aliasLabel, aliasId).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "View User Response: %v\n", res)
	fmt.Fprintf(os.Stdout, "User: %v\n", user)
}

func UpdateUser(aliasLabel string, aliasId string) {
	updates := *onesignal.NewUpdateUserRequestWithDefaults()
	*updates.Properties.Language = "en"

	updatedUser, res, err := apiClient.DefaultApi.UpdateUser(authCtx, appId, aliasLabel, aliasId).UpdateUserRequest(updates).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Update User Response: %v\n", res)
	fmt.Fprintf(os.Stdout, "User: %v\n", updatedUser)
}

func DeleteUser(aliasLabel string, aliasId string) {
	res, err := apiClient.DefaultApi.DeleteUser(authCtx, appId, aliasLabel, aliasId).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Delete User Response: %v\n", res.Body)
}

func ViewUserIdentity(aliasLabel string, aliasId string) {
	identity, res, err := apiClient.DefaultApi.FetchUserIdentity(authCtx, appId, aliasLabel, aliasId).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "View User Identity Response: %v\n", res.Body)
	fmt.Fprintf(os.Stdout, "User Identity: %v\n", identity)
}

func ViewUserIdentityViaSubscription(subscription string) {
	identity, res, err := apiClient.DefaultApi.IdentifyUserBySubscriptionId(authCtx, appId, subscription).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "View User Identity (Subscription) Response: %v\n", res.Body)
	fmt.Fprintf(os.Stdout, "User Identity: %v\n", identity)
}

func CreateAlias(aliasLabel string, aliasId string) {
	alias, res, err := apiClient.DefaultApi.DeleteAlias(authCtx, appId, aliasLabel, aliasId, "").Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Create Alias Response: %v\n", res.Body)
	fmt.Fprintf(os.Stdout, "Alias: %v\n", alias)

}

func CreateAliaViaSubscription() {

}

func DeleteAlias(aliasLabel string, aliasId string, labellToRemove string) {
	alias, res, err := apiClient.DefaultApi.DeleteAlias(authCtx, appId, aliasLabel, aliasId, labellToRemove).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Deleted Alias Response: %v\n", res.Body)
	fmt.Fprintf(os.Stdout, "Alias: %v\n", alias)
}

func CreateSubscription(aliasLabel string, aliasId string) {
	subscription, res, err := apiClient.DefaultApi.CreateSubscription(authCtx, appId, aliasLabel, aliasId).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Create Subscription Response: %v\n", res.Body)
	fmt.Fprintf(os.Stdout, "Alias: %v\n", subscription)
}

func UpdateSubscription(subscriptionId string) {
	res, err := apiClient.DefaultApi.UpdateSubscription(authCtx, appId, subscriptionId).UpdateSubscriptionRequestBody(*onesignal.NewUpdateSubscriptionRequestBodyWithDefaults()).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Update Subscription Response: %v\n", res.Body)
}

func ViewSubscriptionIAMs(subscriptionId string) {
	iams, res, err := apiClient.DefaultApi.GetEligibleIams(authCtx, appId, subscriptionId).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "View Subscription IAMs Response: %v\n", res.Body)
	fmt.Fprintf(os.Stdout, "In-App Messages: %v\n", iams)
}

func DeleteSubscription(subscriptionId string) {
	res, err := apiClient.DefaultApi.DeleteSubscription(authCtx, appId, subscriptionId).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Delete Subscription Response: %v\n", res.Body)
}

func TransferSubscription(subscriptionId string) {
	transferredSubscription, res, err := apiClient.DefaultApi.TransferSubscription(authCtx, appId, subscriptionId).TransferSubscriptionRequestBody(*onesignal.NewTransferSubscriptionRequestBodyWithDefaults()).Execute()

	if err != nil {
		WriteError(err)
		return
	}

	fmt.Fprintf(os.Stdout, "Transfer Subscription Response: %v\n", res.Body)
	fmt.Fprintf(os.Stdout, "Transferred Subscription: %v\n", transferredSubscription)
}
