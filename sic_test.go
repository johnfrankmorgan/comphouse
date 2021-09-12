package comphouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSICDescription(t *testing.T) {
	assert := assert.New(t)

	assert.NotEqual("Unknown", SIC("1010").Description())
}

func TestSICDescriptionNotFound(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("Unknown", SIC("0000").Description())
}
