package comphouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusCodeToError(t *testing.T) {
	type test struct {
		inp  int
		exp  error
		name string
	}

	tests := []test{
		{401, ErrUnauthorized, "unauthorized"},
		{404, ErrNotFound, "not found"},
		{429, ErrTooManyRequests, "too many requests"},
		{300, ErrUnexpectedStatus, "multiple choices"},
		{400, ErrUnexpectedStatus, "bad request"},
		{500, ErrUnexpectedStatus, "server error"},
		{200, nil, "ok"},
		{201, nil, "created"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			assert.Equal(test.exp, statusCodeToError(test.inp))
		})
	}
}
