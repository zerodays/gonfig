package ini

import (
	"gopkg.in/ini.v1"
)

type Provider struct {
	file *ini.File
}

// New returns new ini provider instance with specified ini config source.
func New(source interface{}) (*Provider, error) {
	file, err := ini.Load(source)
	if err != nil {
		return nil, err
	}

	return &Provider{
		file: file,
	}, nil
}

// GetString gets string from ini file for specified parameters.
func (p *Provider) GetString(_, section, name string) (string, error) {
	return p.file.Section(section).Key(name).String(), nil
}

// GetInt gets int from ini file for specified parameters.
func (p *Provider) GetInt(_, section, name string) (int, error) {
	return p.file.Section(section).Key(name).Int()
}

// GetBool gets bool from ini file for specified parameters.
func (p *Provider) GetBool(_, section, name string) (bool, error) {
	return p.file.Section(section).Key(name).Bool()
}

// GetFloat gets float from ini file for specified parameters.
func (p *Provider) GetFloat(_, section, name string) (float64, error) {
	return p.file.Section(section).Key(name).Float64()
}
