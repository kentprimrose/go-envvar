package envvar

import (
	"fmt"
	"os"
	"strings"
)

// Getval gets the value of a variable from env (or defaults).
func Getval(valName string, opts ...string) (string, error) {

	// Get the result right away if it's there.
	result := os.Getenv(valName)
	if result != "" {
		return result, nil
	}

	// Check whether the empty string really IS the value.
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == valName {
			return result, nil
		}
	}

	// If a default value was provided, return it.
	if len(opts) > 0 {
		defaultValue := opts[0]
		return defaultValue, nil
	}

	// Return an error to indicate nothing found.
	return result, fmt.Errorf("Environment variable %s not found", valName)
}
