package comphouse

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkCompanyEndpointHandlesError(t *testing.T, f func(*CompanyEndpoint) error) {
	assert := assert.New(t)

	ts, c := createTestServer(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(400)
	})

	defer ts.Close()

	err := f(c.Company(EnglishCompanyNo(1)))
	assert.Error(err)
}

func TestCompanyEndpointHandlesErrors(t *testing.T) {
	type test struct {
		name string
		f    func(*CompanyEndpoint) error
	}

	tests := []test{
		{
			"CompanyEndpoint.Profile",
			func(c *CompanyEndpoint) error {
				_, err := c.Profile()
				return err
			},
		},
		{
			"CompanyEndpoint.RegisteredOfficeAddress",
			func(c *CompanyEndpoint) error {
				_, err := c.RegisteredOfficeAddress()
				return err
			},
		},
		{
			"CompanyEndpoint.Officers",
			func(c *CompanyEndpoint) error {
				_, err := c.Officers()
				return err
			},
		},
		{
			"CompanyEndpoint.Appointments",
			func(c *CompanyEndpoint) error {
				_, err := c.Appointments("")
				return err
			},
		},
		{
			"CompanyEndpoint.Registers",
			func(c *CompanyEndpoint) error {
				_, err := c.Registers()
				return err
			},
		},
		{
			"CompanyEndpoint.Charges",
			func(c *CompanyEndpoint) error {
				_, err := c.Charges()
				return err
			},
		},
		{
			"CompanyEndpoint.Charge",
			func(c *CompanyEndpoint) error {
				_, err := c.Charge("")
				return err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkCompanyEndpointHandlesError(t, test.f)
		})
	}
}

func TestCompanyEndpointDecodesResponses(t *testing.T) {
	type test struct {
		name string
		f    func(c *CompanyEndpoint) (interface{}, error)
	}

	tests := []test{
		{
			"CompanyEndpoint.Profile",
			func(c *CompanyEndpoint) (interface{}, error) {
				return c.Profile()
			},
		},
		{
			"CompanyEndpoint.RegisteredOfficeAddress",
			func(c *CompanyEndpoint) (interface{}, error) {
				return c.RegisteredOfficeAddress()
			},
		},
		{
			"CompanyEndpoint.Officers",
			func(c *CompanyEndpoint) (interface{}, error) {
				return c.Officers()
			},
		},
		{
			"CompanyEndpoint.Appointments",
			func(c *CompanyEndpoint) (interface{}, error) {
				return c.Appointments("")
			},
		},
		{
			"CompanyEndpoint.Registers",
			func(c *CompanyEndpoint) (interface{}, error) {
				return c.Registers()
			},
		},
		{
			"CompanyEndpoint.Charges",
			func(c *CompanyEndpoint) (interface{}, error) {
				return c.Charges()
			},
		},
		{
			"CompanyEndpoint.Charge",
			func(c *CompanyEndpoint) (interface{}, error) {
				return c.Charge("")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			ts, c := createTestServerWithResponse(map[string]string{"company_name": "Company Name"})

			defer ts.Close()

			got, err := test.f(c.Company(EnglishCompanyNo(1)))

			assert.NoError(err)
			assert.NotNil(got)
		})
	}
}
