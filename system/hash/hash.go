package hash

import (
	"golang.org/x/crypto/bcrypt"
)

// hashing algorithm and then return the hashed password
// as a hex string
func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashed)
}

// Check if two passwords match
func DoPasswordsMatch(strhashedPassword, currPassword string) bool {
	hashedPassword := []byte(strhashedPassword)
	password := []byte(currPassword)

	success := bcrypt.CompareHashAndPassword(hashedPassword, password)

	if success == nil {
		return true
	} else {
		return false
	}
}
