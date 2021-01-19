package multauth

import (
	"strings"
	"errors"
)

type UserPhoneService struct {
	UserService
	Provider UserServiceProviderInterface `db:"-"`

	Phone string `db:"phone" json:"phone"`
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

	if service.Phone == "" {
		return errors.New("Error")
	}

	passcode, err := generateOTP(length)
	if err != nil {
		return errors.New("Error")
	}

	errSend := service.Provider.Send(service.Phone, passcode)
	if errSend != nil { return errSend }

	service.Passcode = passcode
	return nil
}

func (service UserPhoneService) CheckPasscode(value string) bool {
	return service.Passcode == value
}
