package helpers

import "golang.org/x/crypto/bcrypt"

// HashBcrypt hash string to bcrypt string
func HashBcrypt(input string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(input), 14)
	return string(bytes), err
}
