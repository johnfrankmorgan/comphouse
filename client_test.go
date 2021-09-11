package comphouse

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClientRequestInvalidRequest(t *testing.T) {
	assert := assert.New(t)

	c := NewClient(Config{})
	req, err := c.request("/")

	assert.Nil(req)
	assert.Error(err)
	assert.Contains(err.Error(), "invalid method")
}

func TestClientGetAuthenticationFails(t *testing.T) {
	assert := assert.New(t)
	expErr := errors.New("auth failed")

	c := NewClient(Config{
		Auth: TestAuthenticator(func(_ *http.Request) error { return expErr }),
		Host: "localhost",
	})

	resp, err := c.get()

	assert.Nil(resp)
	assert.Error(err)
	assert.Same(expErr, err)
}

func TestClientCompanyFailingRequest(t *testing.T) {
	assert := assert.New(t)

	c := NewClient(Config{
		Host: "<>",
	})

	company, err := c.Company(EnglishCompanyNo(1))

	assert.Nil(company)
	assert.Error(err)
	assert.IsType(&url.Error{}, err)
}

func TestClientCompanyInvalidStatusCode(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.WriteHeader(404)
		}),
	)

	defer ts.Close()

	c := NewClient(Config{
		Host: ts.Listener.Addr().String(),
	})
	c.proto = "http"

	company, err := c.Company(EnglishCompanyNo(1))

	assert.Nil(company)
	assert.Error(err)
	assert.Same(ErrNotFound, err)
}

func TestClientCompanySuccessful(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			fmt.Fprintf(w, `{"company_name": "Company Name"}`)
		}),
	)

	defer ts.Close()

	c := NewClient(Config{
		Host: ts.Listener.Addr().String(),
	})
	c.proto = "http"

	company, err := c.Company(EnglishCompanyNo(1))

	assert.NoError(err)
	assert.Equal("Company Name", company.Name)
}
