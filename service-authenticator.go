package multauth

import (
	"errors"
	"github.com/pquerna/otp/totp"
)

type UserAuthenticatorService struct {
	UserService
	Key string `db:"key" json:"key"`
}

func (service *UserAuthenticatorService) Init(data map[string]interface{}) error {
	issuer, issuerOk := data["Issuer"].(string)
	if !issuerOk {
		issuer = DEFAULT_SERVICE_ISSUER
	}

	account, accountOk := data["Account"].(string)
	if !accountOk {
		account = DEFAULT_SERVICE_ACCOUNT
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: account,
	})

	if err != nil {
		return errors.New("Error")
	}

	service.Key = key.Secret()
	return nil
}

func (service UserAuthenticatorService) CheckPasscode(value string) bool {
	return totp.Validate(value, service.Key)
}
