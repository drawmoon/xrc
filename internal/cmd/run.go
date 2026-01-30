package cmd

import (
	"flag"
	"os"
)

// Run handles the "run" command.
func Run(fs *flag.FlagSet) {
	fs.Parse(os.Args[2:])
}
