package multauth

import (
	"errors"
	"github.com/pquerna/otp/totp"
)

type UserPhoneService struct {
	UserService
}
