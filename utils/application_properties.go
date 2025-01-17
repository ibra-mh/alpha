
package utils

import (
        "os"
)

// GetProperty retrieves the value of an environment variable.
// It returns an empty string if the variable is not set.
func GetProperty(key string) string {
        return os.Getenv(key)
}