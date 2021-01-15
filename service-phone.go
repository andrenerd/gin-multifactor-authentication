package multauth

import (
	_ "errors"
	_ "github.com/pquerna/otp/totp"
)

type UserPhoneService struct {
	UserService
}
