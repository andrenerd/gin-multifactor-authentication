package multauth

import (
	"errors"
	"strings"
	"net/url"
	"net/http"
	"encoding/json"
)

const (
)

var client = &http.Client{}

type UserVonageServiceProvider struct {
}

// todo: make as goroutine?
func (provider UserVonageServiceProvider) Send(to string, message string) error {
	return nil
}
