package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword Hashes a password using bcrypt
func HashPassword(password string) (string, error) {
	// Generate a salt with a cost of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// ComparePasswords Compares a hashed password with a plaintext password
func ComparePasswords(hashedPassword, enteredPassword string) error {
	// Compare the hashed password with the plaintext password
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(enteredPassword))
}
