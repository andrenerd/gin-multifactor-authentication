package multauth

import (
	_ "fmt"
	"strings"
	"errors"
)

type UserPhoneService struct {
	UserService
	Passcode string `db:"passcode" json:"passcode"`
}

func (service *UserPhoneService) Init(data map[string]interface{}) error {
	length, lengthOk := data["Length"].(int)
	if !lengthOk {
		length = DEFAULT_SERVICE_LENGTH
	}

	service.Passcode = strings.Repeat("0", length) // experimental
	return nil
}

func (service *UserPhoneService) SetPasscode() error {
	length := len(service.Passcode)
	if length == 0 {
		length = DEFAULT_SERVICE_LENGTH
	}

	passcode, err := generateOTP(length)
	if err != nil {
		return errors.New("Error")
	}

	service.Passcode = passcode
	return nil
}