package comphouse

import (
	"bytes"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// SearchEndpoint is a struct that can be used to perform searches using the
// Companies House REST API
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/search
type SearchEndpoint struct {
	Client *Client
}

// SearchParams are used as input to the various search functions
type SearchParams struct {
	Query        string
	ItemsPerPage int
	StartIndex   int
}

// Encode converts SearchParams into an escaped string suitable for use in
// query strings
func (m SearchParams) Encode() string {
	var buff bytes.Buffer

	buff.WriteString("q=")
	buff.WriteString(url.QueryEscape(m.Query))

	if m.ItemsPerPage > 0 {
		buff.WriteString("&items_per_page=")
		buff.WriteString(url.QueryEscape(strconv.Itoa(m.ItemsPerPage)))
	}

	if m.StartIndex > 0 {
		buff.WriteString("&start_index=")
		buff.WriteString(url.QueryEscape(strconv.Itoa(m.StartIndex)))
	}

	return buff.String()
}

// helper function to format search params into a search path
func (m *SearchEndpoint) path(params SearchParams, extra ...string) string {
	p := "/search"

	if len(extra) > 0 {
		p += "/" + strings.Join(extra, "/")
	}

	return fmt.Sprintf("%s?%s", p, params.Encode())
}

// Search companies, officers and disqualified officers
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/search/search-all
func (m *SearchEndpoint) All(params SearchParams) (*Search, error) {
	s := &Search{}

	if err := m.Client.GetJSON(m.path(params), s); err != nil {
		return nil, err
	}

	return s, nil
}

// Search company information
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/search/search-companies
func (m *SearchEndpoint) Companies(params SearchParams) (*CompanySearch, error) {
	s := &CompanySearch{}

	if err := m.Client.GetJSON(m.path(params, "companies"), s); err != nil {
		return nil, err
	}

	return s, nil
}

// Search for officer information
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/search/search-officers
func (m *SearchEndpoint) Officers(params SearchParams) (*OfficerSearch, error) {
	s := &OfficerSearch{}

	if err := m.Client.GetJSON(m.path(params, "officers"), s); err != nil {
		return nil, err
	}

	return s, nil
}

// Search for disqualified officer information
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/search/search-disqualified-officers
func (m *SearchEndpoint) DisqualifiedOfficers(params SearchParams) (*DisqualifiedOfficerSearch, error) {
	s := &DisqualifiedOfficerSearch{}

	if err := m.Client.GetJSON(m.path(params, "disqualified-officers"), s); err != nil {
		return nil, err
	}

	return s, nil
}
