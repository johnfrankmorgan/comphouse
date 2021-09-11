package comphouse

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Default values used when creating a new Client
const (
	DefaultProtocol = "https"
	DefaultTimeout  = time.Second * 30
)

// Client is a http.Client wrapper to make interacting with the Companies
// House API easier
type Client struct {
	Auth     Authenticator
	Host     string
	Protocol string
	HTTP     *http.Client
}

// NewClient creates a new Client for the specified host. Requests will be
// authenticated using the provided Authenticator
func NewClient(host string, auth Authenticator) *Client {
	if auth == nil {
		auth = APIKey("")
	}

	return &Client{
		Auth:     auth,
		Host:     host,
		Protocol: DefaultProtocol,
		HTTP: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

// URL returns a formatted URL for the Client's configured protocol and host
func (m *Client) URL(path string) string {
	path = strings.TrimPrefix(path, "/")
	return fmt.Sprintf("%s://%s/%s", m.Protocol, m.Host, path)
}

// NewRequest is a helper method to create a new authenticated HTTP request
func (m *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, m.URL(path), body)
	if err != nil {
		return nil, err
	}

	if err := m.Auth.Authenticate(req); err != nil {
		return nil, err
	}

	return req, nil
}

// Do creates a new request and executes it
func (m *Client) Do(method, path string, body io.Reader) (*http.Response, error) {
	req, err := m.NewRequest(method, path, body)
	if err != nil {
		return nil, err
	}

	resp, err := m.HTTP.Do(req)
	if err != nil {
		return nil, err
	}

	if err := statusCodeToError(resp.StatusCode); err != nil {
		return nil, err
	}

	return resp, nil
}

// Get performs a simple GET request to the specified path
func (m *Client) Get(path string) (*http.Response, error) {
	return m.Do(http.MethodGet, path, nil)
}

// GetJSON performs a GET request to the specified path and attempts to decode
// the response into the passed interface
func (m *Client) GetJSON(path string, dest interface{}) error {
	resp, err := m.Get(path)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(dest)
}

// Company creates a new CompanyEndpoint that can be used to fetch company
// information
func (m *Client) Company(companyNo CompanyNumber) *CompanyEndpoint {
	return &CompanyEndpoint{Client: m, Number: companyNo}
}
