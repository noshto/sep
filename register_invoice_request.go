package sep

import (
	"encoding/xml"
	"fmt"
	"time"
)

// RegisterInvoiceRequestEnvelope represents SOAP Envelope of RegisterInvoiceRequest
type RegisterInvoiceRequestEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    Body     `xml:"Body"`
}

// Body represents SOAP Envelope body
type Body struct {
	XSD                    XSDNS                  `xml:"xsd,attr"`
	XSI                    XSINS                  `xml:"xsi,attr"`
	RegisterInvoiceRequest RegisterInvoiceRequest `xml:"RegisterInvoiceRequest"`
}

// XSDNS represents xmlns:xsd namespace
type XSDNS xml.Attr

// XSINS represents xmlns:xsd namespace
type XSINS xml.Attr

// MarshalXMLAttr to conform to Marshaller interface
func (attr *XSDNS) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	fmt.Println(name)
	return xml.Attr{
		Name: xml.Name{
			Space: "xmlns",
			Local: "xsd",
		},
		Value: "http://www.w3.org/2001/XMLSchema",
	}, nil
}

// MarshalXMLAttr to conform to Marshaller interface
func (attr *XSINS) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	fmt.Println(name)
	return xml.Attr{
		Name: xml.Name{
			Space: "xmlns",
			Local: "xsi",
		},
		Value: "http://www.w3.org/2001/XMLSchema-instance",
	}, nil
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
