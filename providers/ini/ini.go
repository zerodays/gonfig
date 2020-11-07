package ini

import (
	"gopkg.in/ini.v1"
	"strings"
)

type Mode int

const (
	ModeNormal Mode = iota
	ModeUppercase
	ModeLowercase
)

type Provider struct {
	// SectionMode and KeyMode determine weather the section and key parameters
	// should be automatically converted to lower or upper case.
	// Normally SectionMode should be ModeLowercase and KeyMode should be ModeUppercase.
	SectionMode, KeyMode Mode

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

// key returns Key for specified section and name.
func (p *Provider) key(section, name string) *ini.Key {
	// Convert section to lower or upper case based on mode.
	switch p.SectionMode {
	case ModeUppercase:
		section = strings.ToUpper(section)
	case ModeLowercase:
		section = strings.ToLower(section)
	}

	// Convert name to lower or upper case based on mode.
	switch p.KeyMode {
	case ModeUppercase:
		name = strings.ToUpper(name)
	case ModeLowercase:
		name = strings.ToLower(name)
	}

	return p.file.Section(section).Key(name)
}

// GetString gets string from ini file for specified parameters.
func (p *Provider) GetString(_, section, name string) (string, error) {
	return p.key(section, name).String(), nil
}

// GetInt gets int from ini file for specified parameters.
func (p *Provider) GetInt(_, section, name string) (int, error) {
	return p.key(section, name).Int()
}

// GetBool gets bool from ini file for specified parameters.
func (p *Provider) GetBool(_, section, name string) (bool, error) {
	return p.key(section, name).Bool()
}

// GetFloat gets float from ini file for specified parameters.
func (p *Provider) GetFloat(_, section, name string) (float64, error) {
	return p.key(section, name).Float64()
}
