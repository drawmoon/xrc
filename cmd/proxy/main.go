package proxy

import (
	"flag"
	"fmt"
	"os"

	"github.com/drawmoon/xrc/internal/cmd"
	"github.com/drawmoon/xrc/internal/config"
	"github.com/drawmoon/xrc/internal/logger"
)

const (
	MSG_MAIN_CMD_USAGE = "Usage: xrc [run|update|subscription]"
)

func main() {
	// Load configuration
	c, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	var updateTarget string

	// Run command flags
	runFs := flag.NewFlagSet("run", flag.ExitOnError)
	runFs.BoolVar(&c.Verbose, "v", false, "Verbose logs")
	runFs.IntVar(&c.SocksPort, "socks", 7897, "Socks proxy port")
	runFs.IntVar(&c.HttpPort, "http", 7897, "HTTP proxy port")

	// Update command flags
	updateFs := flag.NewFlagSet("update", flag.ExitOnError)
	updateFs.StringVar(&updateTarget, "type", "app", "Update type: app|subscription|core|geodata")

	// Subscription command flags
	subscriptionFs := flag.NewFlagSet("subscription", flag.ExitOnError)

	if len(os.Args) < 2 {
		fmt.Println(MSG_MAIN_CMD_USAGE)
		os.Exit(1)
	}

	// Initialize logger based on verbosity
	logger.Setup(c.Verbose)

	switch os.Args[1] {
	case "run":
		cmd.Run(runFs)
	case "update":
		cmd.Update(updateFs, updateTarget)
	case "subscription":
		cmd.Subscription(subscriptionFs)
	default:
		fmt.Println(MSG_MAIN_CMD_USAGE)
		os.Exit(1)
	}
}
