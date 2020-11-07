// Provider that gets its values from environment variables.

package environment

import (
	"fmt"
	"github.com/zerodays/gonfig"
	"os"
	"strconv"
	"strings"
)

type Provider struct{}

// generateKey generates key for environment variable with specified parameters.
func generateKey(appName, section, name string) string {
	key := fmt.Sprintf("%s_%s_%s", appName, section, name)
	return strings.ToUpper(key)
}

// GetString gets string from environment variable for specified parameters.
func (p Provider) GetString(appName, section, name string) (string, error) {
	val := os.Getenv(generateKey(appName, section, name))
	if val == "" {
		return "", gonfig.ErrNoValue
	}

	return val, nil
}

// GetInt gets int from environment variable for specified parameters.
func (p Provider) GetInt(appName, section, name string) (int, error) {
	val, err := p.GetString(appName, section, name)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(val, 10, 64)
	return int(i), err
}

// GetBool gets bool from environment variable for specified parameters.
func (p Provider) GetBool(appName, section, name string) (bool, error) {
	val, err := p.GetString(appName, section, name)
	if err != nil {
		return false, err
	}

	return strconv.ParseBool(val)
}

// GetFloat gets float from environment variable for specified parameters.
func (p Provider) GetFloat(appName, section, name string) (float64, error) {
	val, err := p.GetString(appName, section, name)
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(val, 64)
}
