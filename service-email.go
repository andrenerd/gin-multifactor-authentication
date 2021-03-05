package multauth

import (
	"strings"
	"errors"
)

type UserEmailService struct {
	UserService
	Provider UserServiceProviderInterface `db:"-"`

	Email string `db:"email" json:"email"`
	Passcode string `db:"passcode" json:"passcode"`
}

func (service *UserEmailService) Init(data map[string]interface{}) error {
	length, lengthOk := data["Length"].(int)
	if !lengthOk {
		length = DEFAULT_SERVICE_LENGTH
	}

	service.Passcode = strings.Repeat("0", length) // experimental
	return nil
}

func (service *UserEmailService) SetPasscode() error {
	length := len(service.Passcode)
	if length == 0 {
		length = DEFAULT_SERVICE_LENGTH
	}

	if service.Email == "" {
		return errors.New("Error")
	}

	passcode, err := generateOTP(length)
	if err != nil {
		return errors.New("Error")
	}

	errSend := service.Provider.Send(service.Email, passcode)
	if errSend != nil { return errSend }

	service.Passcode = passcode
	return nil
}

func (service UserEmailService) CheckPasscode(value string) bool {
	return service.Passcode == value
}
