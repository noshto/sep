package sep

// EnvironmentType represents type of Environment
type EnvironmentType string

// EnvironmentType constants
const (
	TEST EnvironmentType = "TEST"
	PROD EnvironmentType = "PROD"
)

// Client represents Buyer, but with additional details
type Client struct {
	Name    string `json:"Name"`
	TIN     string `json:"TIN"`
	VAT     string `json:"VAT"`
	Address string `json:"Address"`
	Town    string `json:"Town"`
	Country string `json:"Country"`
}

// Config represents configuration that is used by efi modules
type Config struct {
	// Company details
	Name        string `json:"Name"`
	TIN         string `json:"TIN"`
	VAT         string `json:"VAT"`
	Address     string `json:"Address"`
	Town        string `json:"Town"`
	Country     string `json:"Country"`
	Phone       string `json:"Phone"`
	Fax         string `json:"Fax"`
	BankAccount string `json:"BankAccount"`

	// Environment, eg staging or production
	Environment EnvironmentType `json:"Environment"`

	// EFI-related constants
	OperatorCode string `json:"OperatorCode"`
	TCR          *TCR   `json:"TCR"`
}
