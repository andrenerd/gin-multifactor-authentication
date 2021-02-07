package multauth

import (
	"strings"
	"errors"
)

type UserWhatsappService struct {
	UserService
	Provider UserServiceProviderInterface `db:"-"`

	Whatsapp string `db:"whatsapp" json:"whatsapp"`
	Passcode string `db:"passcode" json:"passcode"`
}

func (service *UserWhatsappService) Init(data map[string]interface{}) error {
	length, lengthOk := data["Length"].(int)
	if !lengthOk {
		length = DEFAULT_SERVICE_LENGTH
	}

	service.Passcode = strings.Repeat("0", length) // experimental
	return nil
}

func (service *UserWhatsappService) SetPasscode() error {
	length := len(service.Passcode)
	if length == 0 {
		length = DEFAULT_SERVICE_LENGTH
	}

	if service.Whatsapp == "" {
		return errors.New("Error")
	}

	passcode, err := generateOTP(length)
	if err != nil {
		return errors.New("Error")
	}

	errSend := service.Provider.Send(service.Whatsapp, passcode)
	if errSend != nil { return errSend }

	service.Passcode = passcode
	return nil
}

func (service UserWhatsappService) CheckPasscode(value string) bool {
	return service.Passcode == value
}
