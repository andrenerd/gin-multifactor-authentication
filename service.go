package multauth

import (
	"crypto/rand"
)

// experimental
const (
	DEFAULT_SERVICE_ISSUER = "Multauth"
	DEFAULT_SERVICE_ACCOUNT = "Me"
	DEFAULT_SERVICE_LENGTH = 6
)

const digits = "1234567890"

func generateOTP(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)

	if err != nil {
		return "", err
	}

	digitsLength := len(digits)
	for i := 0; i < length; i++ {
		buffer[i] = digits[int(buffer[i])%digitsLength]
	}

	return string(buffer), nil
}
