# :house: comphouse

A simple client for the Companies House REST API.

![CI](https://github.com/johnfrankmorgan/comphouse/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/johnfrankmorgan/comphouse/branch/main/graph/badge.svg?token=OXSB3U3QEX)](https://codecov.io/gh/johnfrankmorgan/comphouse)


## Example Usage

```go
func getCompanyName(companyNo int) (string, error) {
    client := comphouse.NewClient(
        "api.company-information.service.gov.uk",
        comphouse.APIKey("my-api-key"),
    )

    client.Hooks.BeforeRequest = append(client.Hooks.BeforeRequest, func (_ *http.Request) {
        fmt.Println("executed before sending a request!")
    })

    client.Hooks.AfterRequest = append(client.Hooks.AfterRequest, func (_ *http.Response) {
        fmt.Println("executed after sending a request!")
    })

    company := client.Company(comphouse.EnglishCompanyNo(companyNo))

    profile, err := company.Profile()
    if err != nil {
        return "", err
    }

    return profile.CompanyName, nil
}
```
