package provider

type Provider interface{}

var providers = make(map[string]Provider)

// AddProvider adds a new provider that
// can be loaded
func AddProvider(name string, provider Provider) error {
	if providers[name] != nil {
		return ErrProviderIsAlreadyRegistered
	}

	providers[name] = provider

	return nil
}

func LoadProvider(name string) (Provider, error) {
	if provider := providers[name]; provider != nil {
		return provider, nil
	}

	return nil, ErrProviderDoesNotExist
}
