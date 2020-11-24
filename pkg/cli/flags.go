package cli

import (
	"flag"
	"os"
)

var args = Arguments{}

// Arguments holds the arguments passed
// to the application
type Arguments struct {
	PlanPath      string
	DisplayHelp   bool
	ListProviders bool
}

func init() {
	// help flags
	flag.BoolVar(&args.DisplayHelp, "help", false, "Displays help")
	flag.BoolVar(&args.ListProviders, "providers", false, "Lists the avalible providers")
}

// ParseArgs parses the arguments passed to
// the application and returns then
func ParseArgs() *Arguments {
	flag.Parse()

	switch {
	case args.DisplayHelp == true:
		PrintHelp()
	case args.ListProviders == true:
		PrintProviderList()
	default:
		if len(os.Args) > 2 {
			args.PlanPath = os.Args[len(os.Args)-1]
		}
	}

	return &args
}

// Validate checks the arguments passed to the
// application are valid which includes ensuring
// there's a value for PlanPath
func (a *Arguments) Validate() (*Arguments, error) {
	if !a.DisplayHelp && !a.ListProviders && a.PlanPath == "" {
		return nil, ErrMissingPlanPath
	}

	return a, nil
}
