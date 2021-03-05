package multauth

import (
	"net/smtp"
)

const (
	// reserved
)

type UserSmtpServiceProvider struct {
	Host string
	Port string
	From string
	Password string
}

func (provider UserSmtpServiceProvider) Send(to string, message string) error {
	auth := smtp.PlainAuth("", provider.From, provider.Password, provider.Host)

	err := smtp.SendMail(
		provider.Host + ":" + provider.Port,
		auth,
		provider.From,
		[]string{to},
		[]byte(message),
	)

	if (err == nil) {
		return nil
	} else {
		return err // experimental
	}
}
