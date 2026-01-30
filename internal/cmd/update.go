package cmd

import (
	"flag"
	"os"
)

// Update handles the "update" command.
func Update(fs *flag.FlagSet, updateTarget string) {
	fs.Parse(os.Args[2:])
}
