package sep

import (
	"encoding/xml"
	"time"
)

// RegisterCashDepositRequest represents complete details required for registering cash deposit
type RegisterCashDepositRequest struct {
	XMLName     xml.Name `xml:"https://efi.tax.gov.me/fs/schema RegisterCashDepositRequest"`
	ID          string   `xml:"Id,attr"`
	Version     string   `xml:"Version,attr"`
	Header      Header   `xml:"Header"`
	CashDeposit struct {
		CashAmt        Amount                   `xml:"CashAmt,attr"`
		ChangeDateTime DateTime                 `xml:"ChangeDateTime,attr"`
		IssuerTIN      TIN                      `xml:"IssuerTIN,attr"`
		Operation      CashDepositOperationType `xml:"Operation,attr"`
		TCRCode        TCRCode                  `xml:"TCRCode,attr"`
	} `xml:"CashDeposit"`
	Signature Signature `xml:"Signature"`
}

// RegisterCashDepositResponse represents details of server response for RegisterCashDepositRequest
type RegisterCashDepositResponse struct {
	Body struct {
		Fault                       Fault `xml:"Fault,omitempty"`
		RegisterCashDepositResponse struct {
			Header struct {
				UUID         string    `xml:"UUID,attr"`
				RequestUUID  string    `xml:"RequestUUID,attr"`
				SendDateTime time.Time `xml:"SendDateTime,attr"`
			} `xml:"Header,omitempty"`
			FCDC string `xml:"FCDC"`
		} `xml:"RegisterCashDepositResponse,omitempty"`
	} `xml:"Body"`
}
