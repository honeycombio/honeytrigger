package main

import (
	"fmt"
	"net/http"
	"os"

	flag "github.com/jessevdk/go-flags"
)

// BuildID is set by Travis CI
var BuildID string

// UserAgent is what gets included in all http requests to the api
var UserAgent string

type Options struct {
	WriteKey string `short:"k" long:"writekey" description:"Honeycomb write key from https://ui.honeycomb.io/account" required:"true"`
	Dataset  string `short:"d" long:"dataset" description:"Honeycomb dataset name from https://ui.honeycomb.io/dashboard" required:"true"`
	APIHost  string `long:"api_host" hidden:"true" default:"https://api.honeycomb.io/"`
}

var options Options
var parser = flag.NewParser(&options, flag.Default)
var client = http.Client{}
var usage = `-k <writekey> -d <dataset> COMMAND [other flags]

  honeytrigger is a command line utility for manipulating triggers in your
  honeycomb account.

  Writekey and Dataset are both required. Some commands have additional
  arguments.

  'honeytrigger COMMAND --help' will print command-specific flags`

// setVersion sets the internal version ID and updates libhoney's user-agent
func setVersionUserAgent() {
	var version string

	if BuildID == "" {
		version = "dev"
	} else {
		version = BuildID
	}
	UserAgent = fmt.Sprintf("honeytrigger/%s", version)
}

func main() {
	setVersionUserAgent()

	parser.AddCommand("apply", "apply a set of triggers",
		`apply the specified trigger config file.

  This command will check existing triggers and create any that are defined
  in the configuration file but missing (no name match) in the config. It will
  also update existing triggers. It will warn of any triggers that exist but
  that do not match existing configuration file entries.`,
		&ApplyCommand{})

	parser.AddCommand("list", "List all triggers",
		`List all triggers for the specified dataset. 
		
  Trigger definitions will be returned in JSON format.`,
		&ListCommand{})

	// run whichever command is chosen
	parser.Usage = usage
	if _, err := parser.Parse(); err != nil {
		if flagErr, ok := err.(*flag.Error); ok {
			if flagErr.Type == flag.ErrHelp {
				// asking for help isn't a failed run.
				os.Exit(0)
			}
			if flagErr.Type == flag.ErrCommandRequired ||
				flagErr.Type == flag.ErrUnknownFlag ||
				flagErr.Type == flag.ErrRequired {
				fmt.Println("  run 'honeytrigger --help' for full usage details")
			}
		}
		os.Exit(1)
	}
}
