package cli

import (
	"errors"
)

// ErrMissingPlanPath signified no plan path has been passed
var ErrMissingPlanPath = errors.New("miners: Missing Plan path\nUse miner -help for help with command-line options")
