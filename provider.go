package gonfig

// Provider provides a data source for configuration.
type Provider interface {
	GetInt(appName, section, name string) (int, error)
	GetString(appName, section, name string) (string, error)
	GetBool(appName, section, name string) (bool, error)
	GetFloat(appName, section, name string) (float64, error)
}
