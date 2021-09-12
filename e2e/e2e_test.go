package e2e

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/johnfrankmorgan/comphouse"
	"github.com/stretchr/testify/assert"
)

func client(t *testing.T) *comphouse.Client {
	c := comphouse.NewClient(
		"api.company-information.service.gov.uk",
		comphouse.APIKey(os.Getenv("CH_API_KEY")),
	)

	c.Hooks.AfterRequest = append(c.Hooks.AfterRequest, func(resp *http.Response) {
		t.Logf("%s %s %s", resp.Request.Method, resp.Request.URL, resp.Status)
	})

	return c
}

func TestCompanyProfile(t *testing.T) {
	type test struct {
		number comphouse.CompanyNumber
		name   string
		err    error
	}

	tests := []test{
		{comphouse.EnglishCompanyNo(1081551), "ARGOS LIMITED", nil},
		{comphouse.ScottishCompanyNo(311560), "BREWDOG PLC", nil},
		{comphouse.EnglishCompanyNo(1), "", comphouse.ErrNotFound},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%s %s", test.number, test.name), func(t *testing.T) {
			assert := assert.New(t)

			profile, err := client(t).Company(test.number).Profile()

			if test.err != nil {
				assert.Same(test.err, err)
			} else if assert.NoError(err) {
				assert.Equal(test.name, profile.CompanyName)
			}
		})
	}
}

func TestAllResourcesCanBeDecoded(t *testing.T) {
	type test struct {
		name string
		f    func(*comphouse.Client) (interface{}, error)
	}

	companyNumber := comphouse.EnglishCompanyNo(1081551)
	searchParams := comphouse.SearchParams{Query: "argos"}

	tests := []test{
		{
			"CompanyEndpoint.Profile",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Company(companyNumber).Profile()
			},
		},
		{
			"CompanyEndpoint.RegisteredOfficeAddress",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Company(companyNumber).RegisteredOfficeAddress()
			},
		},
		{
			"CompanyEndpoint.Officers",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Company(companyNumber).Officers()
			},
		},
		{
			"CompanyEndpoint.Appointments",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Company(companyNumber).Appointments("41_6e9TvJ63ZibtI8sdNGWvOGoI")
			},
		},
		{
			"CompanyEndpoint.Registers",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Company(companyNumber).Registers()
			},
		},

		{
			"CompanyEndpoint.Charges",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Company(companyNumber).Charges()
			},
		},
		{
			"CompanyEndpoint.Charge",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Company(companyNumber).Charge("n-LzQBYIroD60vcrZtWHICCkqhk")
			},
		},
		{
			"SearchEndpoint.All",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Search().All(searchParams)
			},
		},
		{
			"SearchEndpoint.Companies",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Search().Companies(searchParams)
			},
		},
		{
			"SearchEndpoint.Officers",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Search().Officers(searchParams)
			},
		},
		{
			"SearchEndpoint.DisqualifiedOfficers",
			func(c *comphouse.Client) (interface{}, error) {
				return c.Search().DisqualifiedOfficers(searchParams)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			got, err := test.f(client(t))

			if assert.NoError(err) {
				assert.NotNil(got)
			}
		})
	}
}
