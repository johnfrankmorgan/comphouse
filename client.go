package comphouse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	DefaultTimeout = time.Second * 30
)

type Client struct {
	auth  Authenticator
	host  string
	http  *http.Client
	proto string
}

type Config struct {
	Auth    Authenticator
	Host    string
	Timeout time.Duration
}

func NewClient(cfg Config) *Client {
	if cfg.Timeout == 0 {
		cfg.Timeout = DefaultTimeout
	}

	if cfg.Auth == nil {
		cfg.Auth = APIKey("")
	}

	return &Client{
		auth: cfg.Auth,
		host: cfg.Host,
		http: &http.Client{
			Timeout: cfg.Timeout,
		},
		proto: "https",
	}
}

func (m *Client) url(path ...string) string {
	return fmt.Sprintf("%s://%s/%s", m.proto, m.host, strings.Join(path, "/"))
}

func (m *Client) request(method string, path ...string) (*http.Request, error) {
	req, err := http.NewRequest(method, m.url(path...), nil)
	if err != nil {
		return nil, err
	}
	if err := m.auth.Authenticate(req); err != nil {
		return nil, err
	}
	return req, nil
}

func (m *Client) get(path ...string) (*http.Response, error) {
	req, err := m.request(http.MethodGet, path...)
	if err != nil {
		return nil, err
	}
	return m.http.Do(req)
}

func (m *Client) decode(resp *http.Response, dest interface{}) error {
	if err := statusCodeToError(resp.StatusCode); err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(dest)
}

func (m *Client) Company(id CompanyNumber) (*CompanyProfile, error) {
	resp, err := m.get("company", id.String())
	if err != nil {
		return nil, err
	}
	c := &CompanyProfile{}
	if err := m.decode(resp, c); err != nil {
		return nil, err
	}
	return c, nil
}
