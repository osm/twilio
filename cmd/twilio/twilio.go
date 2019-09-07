package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/osm/twilio"
)

func require(key string) string {
	value := os.Getenv(key)
	if value == "" {
		fmt.Fprintf(os.Stderr, "error: set %s in your environment\n", key)
		os.Exit(1)
	}
	return value
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "usage: %s <to> <message>\n", os.Args[0])
		os.Exit(1)
	}

	accountSID := require("TWILIO_ACCOUNT_SID")
	authToken := require("TWILIO_AUTH_TOKEN")
	from := require("TWILIO_FROM")
	t := twilio.New(accountSID, authToken, from)

	statusCode, err := t.SendSMS(os.Args[1], strings.Join(os.Args[2:], " "))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %w\n", err)
		os.Exit(1)
	}

	fmt.Println(statusCode)
}
