package comphouse

import "strings"

// CompanyEndpoint is a struct that can be used to query the Companies House Public Data API
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference
type CompanyEndpoint struct {
	Client *Client
	Number CompanyNumber
}

// helper method to format a path for a company
func (m *CompanyEndpoint) path(extra ...string) string {
	p := "/company/" + m.Number.String()

	if len(extra) > 0 {
		p += "/" + strings.Join(extra, "/")
	}

	return p
}

// Get the basic company information
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/company-profile/company-profile
func (m *CompanyEndpoint) Profile() (*CompanyProfile, error) {
	c := &CompanyProfile{}

	if err := m.Client.GetJSON(m.path(), c); err != nil {
		return nil, err
	}

	return c, nil
}

// Get the current address of a company
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/registered-office-address/registered-office-address
func (m *CompanyEndpoint) RegisteredOfficeAddress() (*RegisteredOfficeAddress, error) {
	a := &RegisteredOfficeAddress{}

	if err := m.Client.GetJSON(m.path("registered-office-address"), a); err != nil {
		return nil, err
	}

	return a, nil
}

// List of all company officers
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/officers/list
func (m *CompanyEndpoint) Officers() (*OfficerList, error) {
	o := &OfficerList{}

	if err := m.Client.GetJSON(m.path("officers"), o); err != nil {
		return nil, err
	}

	return o, nil
}

// Get details of an individual company officer appointment
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/officers/get-a-company-officer-appointment
func (m *CompanyEndpoint) Appointments(appointmentId string) (*OfficerSummary, error) {
	o := &OfficerSummary{}

	if err := m.Client.GetJSON(m.path("appointments", appointmentId), o); err != nil {
		return nil, err
	}

	return o, nil
}

// Get the company registers information
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/registers/company-registers
func (m *CompanyEndpoint) Registers() (*CompanyRegister, error) {
	r := &CompanyRegister{}

	if err := m.Client.GetJSON(m.path("registers"), r); err != nil {
		return nil, err
	}

	return r, nil
}

// List of charges for a company
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/charges/list
func (m *CompanyEndpoint) Charges() (*ChargeList, error) {
	c := &ChargeList{}

	if err := m.Client.GetJSON(m.path("charges"), c); err != nil {
		return nil, err
	}

	return c, nil
}

// Individual charge information for company
// https://developer-specs.company-information.service.gov.uk/companies-house-public-data-api/reference/charges/get
func (m *CompanyEndpoint) Charge(chargeId string) (*ChargeDetails, error) {
	c := &ChargeDetails{}

	if err := m.Client.GetJSON(m.path("charges", chargeId), c); err != nil {
		return nil, err
	}

	return c, nil
}
