package goenv

import "os"

// LookupEnv adds optional default to os.Lookupenv.
func LookupEnv(valName string, opts ...string) (string, bool) {

	result, found := os.LookupEnv(valName)

	// If result found, return it.
	if found {
		return result, found
	}

	// If a default value was provided, return it.
	if len(opts) > 0 {
		defaultValue := opts[0]
		return defaultValue, found
	}

	// If not found, and no default, just return os.LookupEnv results.
	return result, found
}
