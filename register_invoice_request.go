package sep

import (
	"encoding/xml"
	"time"
)

// RegisterInvoiceRequestEnvelope represents SOAP Envelope of RegisterInvoiceRequest
type RegisterInvoiceRequestEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body     `xml:"Body"`
}

// Body represents SOAP Envelope body
type Body struct {
	XSD                    string                 `xml:"http://www.w3.org/2001/XMLSchema xsd,attr"`
	XSI                    string                 `xml:"http://www.w3.org/2001/XMLSchema-instance xsi,attr"`
	RegisterInvoiceRequest RegisterInvoiceRequest `xml:"RegisterInvoiceRequest"`
}

// RegisterInvoiceRequest represents details neede for registering an invoice
type RegisterInvoiceRequest struct {
	XMLName   xml.Name  `xml:"https://efi.tax.gov.me/fs/schema RegisterInvoiceRequest"`
	ID        string    `xml:"Id,attr"`
	Version   string    `xml:"Version,attr"`
	Header    Header    `xml:"Header"`
	Invoice   Invoice   `xml:"Invoice"`
	Signature Signature `xml:"Signature"`
}

// RegisterInvoiceResponse represents details of server response for RegisterInvoiceRequest
type RegisterInvoiceResponse struct {
	ID      string `xml:"Id,attr,omitempty"`
	Version string `xml:"Version,attr,omitempty"`
	Body    struct {
		Fault                   Fault `xml:"Fault,omitempty"`
		RegisterInvoiceResponse struct {
			Header struct {
				UUID         string    `xml:"UUID,attr"`
				RequestUUID  string    `xml:"RequestUUID,attr"`
				SendDateTime time.Time `xml:"SendDateTime,attr"`
			} `xml:"Header"`
			FIC string `xml:"FIC,omitempty"`
		} `xml:"RegisterInvoiceResponse,omitempty"`
	} `xml:"Body"`
}
