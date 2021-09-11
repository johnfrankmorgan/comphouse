package comphouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnglishCompanyNo(t *testing.T) {
	assert := assert.New(t)

	no := EnglishCompanyNo(100100)

	assert.Equal("00100100", no.String())
	assert.Equal("00100101", no.Next().String())
}

func TestScottishCompanyNo(t *testing.T) {
	assert := assert.New(t)

	no := ScottishCompanyNo(100100)

	assert.Equal("SC100100", no.String())
	assert.Equal("SC100101", no.Next().String())
}
