package comphouse

import "fmt"

type CompanyNumber interface {
	fmt.Stringer
	Next() CompanyNumber
}

var (
	_ CompanyNumber = EnglishCompanyNo(0)
	_ CompanyNumber = ScottishCompanyNo(0)
)

type EnglishCompanyNo uint

func (m EnglishCompanyNo) String() string {
	return fmt.Sprintf("%08d", m)
}

func (m EnglishCompanyNo) Next() CompanyNumber {
	return m + 1
}

type ScottishCompanyNo uint

func (m ScottishCompanyNo) String() string {
	return fmt.Sprintf("SC%06d", m)
}

func (m ScottishCompanyNo) Next() CompanyNumber {
	return m + 1
}
