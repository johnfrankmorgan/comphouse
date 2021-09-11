# :house: comphouse

A simple client for the Companies House REST API.


## Example Usage

```go
func getCompanyName(companyNo int) (string, error) {
	client := comphouse.NewClient(
		"api.company-information.service.gov.uk",
		comphouse.APIKey("my-api-key"),
	)

	company := client.Company(comphouse.EnglishCompanyNo(companyNo))

	profile, err := company.Profile()
	if err != nil {
		return "", err
	}

	return profile.CompanyName, nil
}
```
