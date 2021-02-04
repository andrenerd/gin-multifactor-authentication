package multauth

import (
	"errors"
	"strings"
	"net/smtp"
)

const (
	// URL_BASE = "https://api.twilio.com/2010-04-01/Accounts/"
	// URL_MESSAGES = "/Messages.json"
	// HEADER_CONTENT_TYPE = "application/x-www-form-urlencoded"
	// HEADER_ACCEPT = "application/json"
)

type UserSmtpServiceProvider struct {
	Host string
	Port string
	// From string
	// Password string
}

func (provider UserSmtpServiceProvider) Send(to string, message string) error {
	return nil
}
