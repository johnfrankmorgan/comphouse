package comphouse

import (
	"fmt"
	"strconv"
	"strings"
)

// CompanyNumber is an interface to convert a value into a valid company number
type CompanyNumber interface {
	fmt.Stringer
}

// CompanyNumberFromString creates a new CompanyNumber from the provided string
// It returns an error if the input cannot be converted to a CompanyNumber
func CompanyNumberFromString(s string) (CompanyNumber, error) {
	s = strings.ToUpper(s)

	f := func(num uint64) CompanyNumber {
		return EnglishCompanyNo(num)
	}

	if strings.HasPrefix(s, "SC") {
		f = func(num uint64) CompanyNumber {
			return ScottishCompanyNo(num)
		}
		s = s[2:]
	}

	num, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return nil, err
	}

	return f(num), nil
}

// EnglishCompanyNo is a CompanyNumber implementation for English companies
type EnglishCompanyNo uint

// String satisfies the CompanyNumber interface
func (m EnglishCompanyNo) String() string {
	return fmt.Sprintf("%08d", m)
}

// ScottishCompanyNo is a CompanyNumber implementation for Scottish companies
type ScottishCompanyNo uint

// String satisfies the CompanyNumber interface
func (m ScottishCompanyNo) String() string {
	return fmt.Sprintf("SC%06d", m)
}
