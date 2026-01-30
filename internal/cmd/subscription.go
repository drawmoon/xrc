package cmd

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
)

const (
	MSG_SUBSCRIPTION_CMD_USAGE = "Usage: xrc subscription [list|set-url|remove]"
)

// Subscription handles the "subscription" command.
func Subscription(fs *flag.FlagSet) {
	fs.Parse(os.Args[2:])

	remainingArgs := fs.Args()
	if len(remainingArgs) < 1 {
		fmt.Println(MSG_SUBSCRIPTION_CMD_USAGE)
		os.Exit(1)
	}

	switch remainingArgs[0] {
	case "list":
		// List subscriptions
	case "set-url":
		// Set subscription URL
		if len(remainingArgs) < 2 {
			slog.Debug("Error: set-url requires a URL\nHint: xrc subscription set-url \"http://...\"")
			return
		}
		url := remainingArgs[1]
		slog.Debug("Setting subscription URL", "url", url)
	case "remove":
		// Remove subscription
	default:
		fmt.Println(MSG_SUBSCRIPTION_CMD_USAGE)
		os.Exit(1)
	}
}
