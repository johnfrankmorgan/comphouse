package comphouse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnglishCompanyNo(t *testing.T) {
	assert := assert.New(t)

	no := EnglishCompanyNo(100100)

	assert.Equal("00100100", no.String())
}

func TestScottishCompanyNo(t *testing.T) {
	assert := assert.New(t)

	no := ScottishCompanyNo(100100)

	assert.Equal("SC100100", no.String())
}

func TestCompanyNumberFromString(t *testing.T) {
	type test struct {
		inp string
		exp CompanyNumber
	}

	tests := []test{
		{"SC000104", ScottishCompanyNo(104)},
		{"10010010", EnglishCompanyNo(10010010)},
		{"SC123123", ScottishCompanyNo(123123)},
		{"1", EnglishCompanyNo(1)},
	}

	for _, test := range tests {
		t.Run(test.inp, func(t *testing.T) {
			assert := assert.New(t)

			number, err := CompanyNumberFromString(test.inp)

			if assert.NoError(err) {
				assert.Equal(test.exp, number)
			}
		})
	}

	t.Run("handles errors", func(t *testing.T) {
		assert := assert.New(t)

		number, err := CompanyNumberFromString("SC")
		if assert.Error(err) {
			assert.Nil(number)
			assert.Contains(err.Error(), "invalid syntax")
		}
	})
}
