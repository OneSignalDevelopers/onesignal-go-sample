![OneSignal](https://raw.githubusercontent.com/OneSignalDevelopers/.github/main/assets/onesignal-banner.png)

<div align="center">
  <a href="https://documentation.onesignal.com/docs/onboarding-with-onesignal" target="_blank">Quickstart</a>
  <span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
  <a href="https://onesignal.com/" target="_blank">Website</a>
  <span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
  <a href="https://documentation.onesignal.com/docs" target="_blank">Docs</a>
  <span>&nbsp;&nbsp;•&nbsp;&nbsp;</span>
  <a href="https://github.com/OneSignalDevelopers" target="_blank">Examples</a>
  <br />
  <hr />
</div>

# OneSignal Go Client Sample

OneSignal makes engaging customers simple and is the fastest, most reliable service to send push notifications, in-app messages, SMS, and emails.

This project demonstrates sending push notifications using the [onesignal-go-api](https://pkg.go.dev/github.com/OneSignal/onesignal-go-api) Go package. You can use this project as a boilerplate or reference to start your project.

## 🚦 Getting started

This project assumes that you already have a OneSignal app created with push notifications setup. If you don't yet have a OneSignal app, [create one](https://documentation.onesignal.com/docs/apps-organizations#create-an-app) first, then follow the steps below to integrate the OneSignal SDK into your [Android](https://documentation.onesignal.com/docs/android-sdk-setup) or [iOS](https://documentation.onesignal.com/docs/ios-sdk-setup) app.

1. Rename _.env.example_ to _.env_, `mv .env.example .env`
2. Set environment variables
3. Run the app, `go run .`

Running this example will result in a notification being sent to **all subscribed** users of your OneSignal app and the following output:

```shell
╭─iamwill@kronos ~/code/@onesignalDevelopers/onesignal-go-sample ‹main●›
╰─$ go run .
Response from `CreateNotification`: &{fa7b2782-3ef7-4f44-b8f3-bba3721797a5 5 <nil> <nil> map[]}
Notification ID: fa7b2782-3ef7-4f44-b8f3-bba3721797a5
```

### Environment variables

| Name          | Where to find it                                                                |
| :------------ | :------------------------------------------------------------------------------ |
| APP_ID        | URL in app dashboard -> `https://app.onesignal.com/apps/<APP_ID>`               |
| USER_AUTH_KEY | [OneSignal profile page](https://app.onesignal.com/profile)                     |
| REST_API_KEY  | App settings -> `https://app.onesignal.com/apps/<APP_ID>/settings/keys_and_ids` |

# Support

## Ask a question about OneSignal

You can ask questions about the OneSignal Go API client and related topics in the [onesignal-go-api](https://github.com/onesignal/onesignal-go-api) repository.

🙋‍♂️ [Ask a question](https://github.com/OneSignal/onesignal-go-api/issues/new?assignees=&labels=triage&template=ask-question.yml&title=%5Bquestion%5D%3A+)

## Create a bug report

If you receive an error message or get blocked by an issue, please create a bug report!

🪳 [Create bug report](https://github.com/OneSignal/onesignal-go-api/issues/new?assignees=&labels=bug%2Ctriage&template=bug-report.yml&title=%5BBug%5D%3A+)

# ❤️ Developer Community

For additional resources, please join the [OneSignal Developer Community](https://onesignal.com/onesignal-developers).

Get in touch with us or learn more about OneSignal through the channels below.

- [Follow us on Twitter](https://twitter.com/onesignaldevs) to never miss any updates from the OneSignal team, ecosystem & community
- [Join us on Discord](https://discord.gg/EP7gf6Uz7G) to be a part of the OneSignal Developers community, showcase your work and connect with other OneSignal developers
- [Read the OneSignal Blog](https://onesignal.com/blog/) for the latest announcements, tutorials, in-depth articles & more.
- [Subscribe to us on YouTube](https://www.youtube.com/channel/UCe63d5EDQsSkOov-bIE_8Aw/featured) for walkthroughs, courses, talks, workshops & more.
- [Follow us on Twitch](https://www.twitch.tv/onesignaldevelopers) for live streams, office hours, support & more.

## Show your support

Give a ⭐️ if this project helped you, and watch this repo to stay up to date.
