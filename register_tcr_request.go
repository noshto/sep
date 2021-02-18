package sep

import (
	"encoding/xml"
	"time"
)

// TCR represents details about TCR
type TCR struct {
	Type           TCRType        `xml:"Type,attr"`
	ValidFrom      Date           `xml:"ValidFrom,attr,omitempty"`
	ValidTo        Date           `xml:"ValidTo,attr"`
	TCRIntID       TCRIntID       `xml:"TCRIntID,attr"`
	IssuerTIN      TIN            `xml:"IssuerTIN,attr"`
	SoftCode       SoftCode       `xml:"SoftCode,attr"`
	MaintainerCode MaintainerCode `xml:"MaintainerCode,attr"`
	BusinUnitCode  BusinUnitCode  `xml:"BusinUnitCode,attr"`
}

// RegisterTCRRequest contains details of TCR needed for registering
type RegisterTCRRequest struct {
	XMLName   xml.Name  `xml:"https://efi.tax.gov.me/fs/schema RegisterTCRRequest"`
	ID        string    `xml:"Id,attr"`
	Version   string    `xml:"Version,attr"`
	Header    Header    `xml:"Header"`
	TCR       TCR       `xml:"TCR"`
	Signature Signature `xml:"Signature"`
}

// RegisterTCRResponse represents server response for RegisterTCRRequest
type RegisterTCRResponse struct {
	ID      string `xml:"Id,attr"`
	Version string `xml:"Version,attr"`
	Body    struct {
		Fault               Fault `xml:"Fault,omitempty"`
		RegisterTCRResponse struct {
			Header struct {
				UUID         string    `xml:"UUID,attr"`
				RequestUUID  string    `xml:"RequestUUID,attr"`
				SendDateTime time.Time `xml:"SendDateTime,attr"`
			} `xml:"Header"`
			TCRCode TCRCode `xml:"TCRCode"`
		} `xml:"RegisterTCRResponse,omitempty"`
	} `xml:"Body"`
}
