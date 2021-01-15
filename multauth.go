package multauth

import (
	_ "fmt"
	"errors"
	"reflect"
	"golang.org/x/crypto/bcrypt"
)

type Flow = []string // todo: or drop it?

const (
	PASSWORD = "Password"
	PASSCODE = "Passcode"
	HARDCODE = "Hardcode"
)

type UserInterface interface {
	SetPassword(value string) error
	CheckPassword(value string) bool

	// Should be implemented by app
	Get(fields map[string]interface{}) error // or AuthGet?
	GetUserServices() ([]UserServiceInterface, error) // or AuthGetUserServices? // todo: how about extra params to load specific services?
	Save(fields ...[]string) error // or AuthSave?
}

type User struct {
	Password string `db:"password" json:"password"`
}

func (user *User) SetPassword(value string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(value), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)
	return nil
}

func (user *User) CheckPassword(value string) bool {
	password := reflect.ValueOf(user).Elem().FieldByName(PASSWORD).Interface().(string)
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(value))
	return err == nil
}

func (user *User) Get(fields map[string]interface{}) error {
	return errors.New("Not implemented")
}

func (user *User) GetUserServices() ([]UserServiceInterface, error) {
	return nil, errors.New("Not implemented")
}

func (user *User) Save(fields ...[]string) error {
	return errors.New("Not implemented")
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

type Auth struct {
	Flows []Flow
}

func (auth Auth) CheckPassword(value string, user UserInterface) bool {
	return user.CheckPassword(value)
}

func (auth Auth) CheckPasscode(value string, user UserInterface) bool {
	services, err := user.GetUserServices()
	if err != nil {
		return false
	}

	for _, service := range services {
		if service.CheckPasscode(value) {
			return true
		}
	}

	return false
}

func (auth Auth) CheckHardcode(value string, user UserInterface) bool {
	services, err := user.GetUserServices()
	if err != nil {
		return false
	}

	for _, service := range services {
		if service.CheckHardcode(value) {
			return true
		}
	}

	return false
}

func (auth Auth) checkSecret(secret string, value string, user UserInterface) bool {
	switch secret {
	case PASSWORD:
		return auth.CheckPassword(value, user)

	case PASSCODE:
		return auth.CheckPasscode(value, user)

	case HARDCODE:
		return auth.CheckHardcode(value, user)
	}

	return false
}

func (auth Auth) checkSecrets(secrets []string, data map[string]interface{}, user UserInterface) bool {
	for _, secret := range secrets {
		value := data[secret]
		if value == nil {
			return false
		}

		if !auth.checkSecret(secret, value.(string), user) {
			return false
		}
	}

	return true
}

func (auth Auth) Authenticate(data map[string]interface{}, user UserInterface, flows ...Flow) error {
	var identifier string

	if len(flows) == 0 {
		flows = auth.Flows
	}

	// Get user by identifier
	for _, flow := range flows {
		// Identifier supposed to go first, always
		flowIdentifier := flow[0]

		if v, ok := data[flowIdentifier]; ok {
			if err := user.Get(map[string]interface{}{flowIdentifier: v}); err != nil {
				return errors.New("Error")
			}

			identifier = flowIdentifier
			break
		}
	}

	// Check user secrets
	for _, flow := range flows {
		// Identifier supposed to go first, always
		if identifier == flow[0] {
			if auth.checkSecrets(flow[1:], data, user) {
				return nil
			}
		}
	}

	return errors.New("Error")
}
