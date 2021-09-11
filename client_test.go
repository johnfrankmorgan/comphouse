package comphouse

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createTestServer(f http.HandlerFunc) (*httptest.Server, *Client) {
	ts := httptest.NewServer(f)
	c := NewClient(ts.Listener.Addr().String(), nil)
	c.Protocol = "http"
	return ts, c
}

func TestClientNewRequestInvalidRequest(t *testing.T) {
	assert := assert.New(t)

	req, err := NewClient("localhost", nil).NewRequest("/", "", nil)

	assert.Nil(req)
	assert.Error(err)
	assert.Contains(err.Error(), "invalid method")
}

func TestClientNewRequestAuthenticationFailure(t *testing.T) {
	assert := assert.New(t)
	expErr := errors.New("authentication failed")

	req, err := NewClient("localhost", TestAuthenticator(func(_ *http.Request) error {
		return expErr
	})).NewRequest("GET", "/", nil)

	assert.Nil(req)
	assert.Error(err)
	assert.Same(expErr, err)
}

func TestClientDoInvalidMethod(t *testing.T) {
	assert := assert.New(t)

	resp, err := NewClient("localhost", nil).Do("/", "", nil)

	assert.Nil(resp)
	assert.Error(err)
	assert.Contains(err.Error(), "invalid method")
}

func TestClientDoErrorExecutingRequest(t *testing.T) {
	assert := assert.New(t)

	resp, err := NewClient("<>", nil).Do("GET", "", nil)

	assert.Nil(resp)
	assert.Error(err)
}

func TestClientDoInvalidStatusCode(t *testing.T) {
	assert := assert.New(t)

	ts, c := createTestServer(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(401)
	})

	defer ts.Close()

	resp, err := c.Do("GET", "", nil)

	assert.Nil(resp)
	assert.Error(err)
	assert.Same(ErrUnauthorized, err)
}

func TestClientGetJSONErrorExecutingRequest(t *testing.T) {
	assert := assert.New(t)

	ts, c := createTestServer(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(404)
	})

	defer ts.Close()

	err := c.GetJSON("", nil)

	assert.Error(err)
	assert.Same(ErrNotFound, err)
}

func TestClientGetJSON(t *testing.T) {
	assert := assert.New(t)

	ts, c := createTestServer(func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, `{"name": "Name"}`)
	})

	defer ts.Close()

	var json struct {
		Name string
	}

	err := c.GetJSON("", &json)

	assert.NoError(err)
	assert.Equal("Name", json.Name)
}

func TestClientCompany(t *testing.T) {
	assert := assert.New(t)

	c := NewClient("localhost", nil)

	assert.Same(c, c.Company(nil).Client)
}
