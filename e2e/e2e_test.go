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

func TestSearchAll(t *testing.T) {
	assert := assert.New(t)

	search, err := client(t).Search().All(comphouse.SearchParams{Query: "argos"})

	assert.NoError(err)
	assert.Equal(search.ItemsPerPage, len(search.Items))
}

func TestSearchCompanies(t *testing.T) {
	assert := assert.New(t)

	search, err := client(t).Search().Companies(comphouse.SearchParams{Query: "argos"})

	assert.NoError(err)
	// FIXME: update this once the resources have been added
	assert.NotNil(search)
}

func TestSearchOfficers(t *testing.T) {
	assert := assert.New(t)

	search, err := client(t).Search().Officers(comphouse.SearchParams{Query: "argos"})

	assert.NoError(err)
	// FIXME: update this once the resources have been added
	assert.NotNil(search)
}

func TestSearchDisqualifiedOfficers(t *testing.T) {
	assert := assert.New(t)

	search, err := client(t).Search().DisqualifiedOfficers(comphouse.SearchParams{Query: "argos"})

	assert.NoError(err)
	// FIXME: update this once the resources have been added
	assert.NotNil(search)
}
