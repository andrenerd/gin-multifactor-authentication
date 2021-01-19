package multauth

import (
	"errors"
	"strings"
	"net/url"
	"net/http"
	"encoding/json"
)

const (
	URL_BASE = "https://api.twilio.com/2010-04-01/Accounts/"
	URL_MESSAGES = "/Messages.json"
	HEADER_CONTENT_TYPE = "application/x-www-form-urlencoded"
	HEADER_ACCEPT = "application/json"
)

var client = &http.Client{}

type UserTwilioServiceProvider struct {
	Account string
	Token string
	From string
}

// todo: make as goroutine?
func (provider UserTwilioServiceProvider) Send(to string, message string) error {
	from := provider.From
	if !strings.HasPrefix(from, "+") { from = "+" + from }
	if !strings.HasPrefix(to, "+") { to = "+" + to }

	v := url.Values{}
	v.Set("To", to)
	v.Set("From", from)
	v.Set("Body", message)
	vReader := *strings.NewReader(v.Encode())

	req, _ := http.NewRequest("POST", URL_BASE + provider.Account + URL_MESSAGES, &vReader)
	req.SetBasicAuth(provider.Account, provider.Token)
	req.Header.Add("Content-Type", HEADER_CONTENT_TYPE)
	req.Header.Add("Accept", HEADER_ACCEPT)

	res, _ := client.Do(req)

	if (res.StatusCode >= 200 && res.StatusCode < 300) {
		decoder := json.NewDecoder(res.Body)
		err := decoder.Decode(&map[string]interface{}{})

		if (err == nil) {
			return nil
		} else {
			return err // experimental
		}

	} else {
		return errors.New(res.Status) // experimental
	}
}
