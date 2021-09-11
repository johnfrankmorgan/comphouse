package comphouse

import "net/http"

type Authenticator interface {
	Authenticate(req *http.Request) error
}

type APIKey string

func (m APIKey) Authenticate(req *http.Request) error {
	req.SetBasicAuth(string(m), "")
	return nil
}
