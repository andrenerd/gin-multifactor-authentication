package multauth

import (
	"errors"
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

type UserServiceInterface interface {
	Init(data map[string]interface{}) error
	SetPasscode() error
	CheckPasscode(value string) bool
	SetHardcode(value string) error
	CheckHardcode(value string) bool

	// Should be implemented by app
	// Get() error // or AuthGet?
	Save(fields ...[]string) error // or AuthSave?
}

type UserService struct {
}

func (service *UserService) Init() error {
	return nil
}

func (service *UserService) SetPasscode() error {
	return errors.New("Not implemented")
}

func (service *UserService) CheckPasscode(value string) bool {
	return false
}

func (service *UserService) SetHardcode(value string) error {
	return errors.New("Not implemented")
}

func (service *UserService) CheckHardcode(value string) bool {
	return false
}

func (service *UserService) Verify() error {
	return errors.New("Not implemented")
}

func (service *UserService) Save(fields ...[]string) error {
	return errors.New("Not implemented")
}
