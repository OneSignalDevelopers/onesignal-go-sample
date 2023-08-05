package main

import (
	"onesignal-developers/sample/lib"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	lib.ViewUser("external_id", "test")
}
