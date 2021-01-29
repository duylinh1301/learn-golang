package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

// InterfaceToMap convert data of interface{} to map[string]interface{}
func InterfaceToMap(data interface{}) map[string]interface{} {
	return data.(map[string]interface{})
}

// HashBcrypt hash string to bcrypt string
func HashBcrypt(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(bytes), err
}

// GetValueOrDefault return default value if value nil
func GetValueOrDefault(value interface{}, defaultValue interface{}) interface{} {
	if value == nil || value == "" {
		return defaultValue
	}
	return value
}
