package multauth

const (
	URL_BASE = "https://api.twilio.com/2010-04-01/Accounts/"
	URL_MESSAGES = "/Messages.json"
)

type UserTwilioServiceProvider struct {
	Account string
	Token string
}

// todo: make as goroutine?
func (provider UserTwilioServiceProvider) Send(to string, message string) error {
	return nil
}
