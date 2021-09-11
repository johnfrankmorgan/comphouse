package comphouse

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestAuthenticator func(*http.Request) error

func (m TestAuthenticator) Authenticate(req *http.Request) error {
	return m(req)
}

func TestAPIKeyAuthenticate(t *testing.T) {
	assert := assert.New(t)

	req, err := http.NewRequest(http.MethodGet, "http://localhost", nil)
	assert.NoError(err)

	err = APIKey("test").Authenticate(req)
	assert.NoError(err)

	username, password, ok := req.BasicAuth()

	assert.True(ok)
	assert.Empty(password)
	assert.Equal("test", username)
}
