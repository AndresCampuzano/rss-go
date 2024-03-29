package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the provided HTTP headers and returns it as a string.
// If the API key is not found, an error is returned.
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("API key not found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid API key format")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("invalid API key format")
	}

	return vals[1], nil
}
