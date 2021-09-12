package main

import (
	"fmt"
	"log"
	"os"

	"github.com/johnfrankmorgan/comphouse"
)

func main() {
	c := comphouse.NewClient(
		"api.company-information.service.gov.uk",
		comphouse.APIKey(os.Getenv("CH_API_KEY")),
	)

	results, err := c.Search().Companies(comphouse.SearchParams{
		Query:        "Subway",
		ItemsPerPage: 5,
	})

	if err != nil {
		log.Fatalln(err)
	}

	for _, result := range results.Items {
		companyNo, err := comphouse.CompanyNumberFromString(result.CompanyNumber)
		if err != nil {
			log.Fatalln(err)
		}

		company := c.Company(companyNo)

		profile, err := company.Profile()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("[%s] %s\n", profile.CompanyNumber, profile.CompanyName)
		fmt.Printf("  Address 1:   %s\n", profile.RegisteredOfficeAddress.AddressLine1)
		fmt.Printf("  Town / City: %s\n", profile.RegisteredOfficeAddress.Locality)
		fmt.Printf("  Postcode:    %s\n", profile.RegisteredOfficeAddress.PostalCode)

		fmt.Println("  SIC Codes:")
		for _, sic := range profile.SicCodes {
			fmt.Printf("    [%s] %s\n", sic, sic.Description())
		}

		officers, err := company.Officers()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("  Officers:")
		for _, officer := range officers.Items {
			occupation := officer.Occupation
			if occupation == "" {
				occupation = "Unknown"
			}

			fmt.Printf("    %s (%s)\n", officer.Name, occupation)
		}

		fmt.Println()
	}
}
