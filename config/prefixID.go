package config

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

// Creates a random ID with a custom prefix
// Prefix must be 3 character long
func PrefixID(prefix string) (string, error) {
	const length = 16

	// Check if prefix character length is equal to 3
	if len(prefix) != 3 {
		return "", errors.New("Prefix must be exactly 3 characters long")
	}
	

	b := make([]byte, length)

	// Read size number of bytes into b
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// This prevents slash ( / ) and plus ( + ) which can break URLs
	// URL Encoding replaces these with - (minus) and _ (underscore)
	encoded := base64.URLEncoding.EncodeToString(b)

	prefixedID := fmt.Sprintf("%s_%s", prefix, encoded)

	return prefixedID, nil

	
}