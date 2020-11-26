package provider

import (
	"errors"
)

// ErrProviderIsAlreadyRegistered is thrown when
// a provider has already been registered
var ErrProviderIsAlreadyRegistered = errors.New("Providers cannot be registered multipe times")

// ErrProviderDoesNotExist is thrown when the
// requested provider has not already been registered
var ErrProviderDoesNotExist = errors.New("The requested provider does not exist")
