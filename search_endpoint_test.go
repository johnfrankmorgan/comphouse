package comphouse

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchParamsEncode(t *testing.T) {
	type test struct {
		inp SearchParams
		exp string
	}

	tests := []test{
		{SearchParams{"test@test", 0, 0}, "q=test%40test"},
		{SearchParams{"testing", 100, 0}, "q=testing&items_per_page=100"},
		{SearchParams{"testing", 0, 150}, "q=testing&start_index=150"},
		{SearchParams{"testing", 50, 150}, "q=testing&items_per_page=50&start_index=150"},
	}

	for _, test := range tests {
		t.Run(test.exp, func(t *testing.T) {
			assert := assert.New(t)
			assert.Equal(test.exp, test.inp.Encode())
		})
	}
}

func checkSearchEndpointHandlesError(t *testing.T, f func(*SearchEndpoint) error) {
	assert := assert.New(t)

	ts, c := createTestServer(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(400)
	})

	defer ts.Close()

	err := f(c.Search())
	assert.Error(err)
}

func TestSearchEndpointHandlesErrors(t *testing.T) {
	type test struct {
		name string
		f    func(*SearchEndpoint) error
	}

	tests := []test{
		{
			"SearchEndpoint.All",
			func(s *SearchEndpoint) error {
				_, err := s.All(SearchParams{})
				return err
			},
		},
		{
			"SearchEndpoint.Companies",
			func(s *SearchEndpoint) error {
				_, err := s.Companies(SearchParams{})
				return err
			},
		},
		{
			"SearchEndpoint.Officers",
			func(s *SearchEndpoint) error {
				_, err := s.Officers(SearchParams{})
				return err
			},
		},
		{
			"SearchEndpoint.DisqualifiedOfficers",
			func(s *SearchEndpoint) error {
				_, err := s.DisqualifiedOfficers(SearchParams{})
				return err
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkSearchEndpointHandlesError(t, test.f)
		})
	}
}

func TestSearchEndpointDecodesResponses(t *testing.T) {
	type test struct {
		name string
		f    func(c *SearchEndpoint) (interface{}, error)
	}

	tests := []test{
		{
			"SearchEndpoint.All",
			func(s *SearchEndpoint) (interface{}, error) {
				return s.All(SearchParams{})
			},
		},
		{
			"SearchEndpoint.Companies",
			func(s *SearchEndpoint) (interface{}, error) {
				return s.Companies(SearchParams{})
			},
		},
		{
			"SearchEndpoint.Officers",
			func(s *SearchEndpoint) (interface{}, error) {
				return s.Officers(SearchParams{})
			},
		},
		{
			"SearchEndpoint.DisqualifiedOfficers",
			func(s *SearchEndpoint) (interface{}, error) {
				return s.DisqualifiedOfficers(SearchParams{})
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert := assert.New(t)

			ts, c := createTestServerWithResponse(map[string]string{"company_name": "Company Name"})

			defer ts.Close()

			got, err := test.f(c.Search())

			assert.NoError(err)
			assert.NotNil(got)
		})
	}
}
