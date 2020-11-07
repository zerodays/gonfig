package gonfig

import "errors"

var ErrNoValue = errors.New("no value")

type Config struct {
	AppName string

	providers []Provider
}

// New creates new config instance with specified providers.
// Providers are ordered from lowest to highest priority.
func New(providers ...Provider) *Config {
	return &Config{
		providers: providers,
	}
}

// AddProvider adds provider to config. Providers
// that are added last have higher priority.
func (cfg *Config) AddProvider(provider Provider) {
	cfg.providers = append(cfg.providers, provider)
}

// GetInt returns integer value for specified section and name.
// If no provider returns a valid value, ErrNoValue is returned.
func (cfg *Config) GetInt(section, name string) (int, error) {
	for i := len(cfg.providers) - 1; i >= 0; i-- {
		val, err := cfg.providers[i].GetInt(cfg.AppName, section, name)
		if err == nil {
			return val, nil
		}
	}

	return 0, ErrNoValue
}

// GetString returns string value for specified section and name.
// If no provider returns a valid value, ErrNoValue is returned.
func (cfg *Config) GetString(section, name string) (string, error) {
	for i := len(cfg.providers) - 1; i >= 0; i-- {
		val, err := cfg.providers[i].GetString(cfg.AppName, section, name)
		if err == nil {
			return val, nil
		}
	}

	return "", ErrNoValue
}

// GetBool returns bool value for specified section and name.
// If no provider returns a valid value, ErrNoValue is returned.
func (cfg *Config) GetBool(section, name string) (bool, error) {
	for i := len(cfg.providers) - 1; i >= 0; i-- {
		val, err := cfg.providers[i].GetBool(cfg.AppName, section, name)
		if err == nil {
			return val, nil
		}
	}

	return false, ErrNoValue
}

// GetFloat returns integer value for specified section and name.
// If no provider returns a valid value, ErrNoValue is returned.
func (cfg *Config) GetFloat(section, name string) (float64, error) {
	for i := len(cfg.providers) - 1; i >= 0; i-- {
		val, err := cfg.providers[i].GetFloat(cfg.AppName, section, name)
		if err == nil {
			return val, nil
		}
	}

	return 0, ErrNoValue
}
