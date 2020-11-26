package plan

import "errors"

// ErrMissingProviderValue is used when the provider
// is missing from the plan
var ErrMissingProviderValue = errors.New("Provider is a required property")
