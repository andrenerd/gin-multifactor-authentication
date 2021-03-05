package multauth

import (
	"errors"
	"bytes"
	"net/url"
	"net/http"
	"encoding/json"
)

const (
	URL_BASE = "https://rest.nexmo.com/sms/json"
)

var client = &http.Client{}

type UserVonageServiceProvider struct {
	Key string
	Secret string
	From string
}

type sendBody struct {
	APIKey string `json:"api_key"`
	APISecret string `json:"api_secret"`
	From string `json:"from"`
	Text string `json:"text"`
	To string `json:"to"`
}

// todo: make as goroutine?
func (provider UserVonageServiceProvider) Send(to string, message string) error {
	body := sendBody{
		APIKey: Key,
		APISecret: Secret,
		From: provider.From,
		To: to,
		Text: message,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	res, err := http.Post(URL_BASE, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	return nil
}
