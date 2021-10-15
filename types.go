package sep

import (
	"encoding/xml"
	"fmt"
	"time"
)

// InvType represents type of invoice
type InvType string

// CountryCode represents country code, 3char
type CountryCode string

// TCRCode represents unique software code
type TCRCode string

// TCRType type of TCR
type TCRType string

// TCRIntID represents TCR internal identifier
type TCRIntID string

// TIN represents Tax Identification Number type
type TIN string

// SoftCode represents uniques software code
type SoftCode string

// MaintainerCode represents uniques software code
type MaintainerCode string

// OperatorCode represents uniques software code
type OperatorCode string

// BusinUnitCode represents type of business unit code
type BusinUnitCode string

// SubseqDelivType - element that says what is the type of subsequent delivery.
// For example, when the issuer sent the message to the CIS, but he did not receive the answer, he can send another message as before, but only with a flag with “Subsequent delivery” note.
// This flag indicates to CIS that the issuer sent another message.
//
// SubseqDelivType - element koji kaže koji je tip naknadnog dostavljanja.
// Na primjer, kada je izdavalac poslao poruku Centralnom registru računa, ali nije primio odgovor iz Centralnog registra računa.
// Može poslati potpuno istu poruku kao i prije, ali samo označeno zastavicom kao „naknadno dostavljanje“.
// Ova zastavica je pokazatelj za CIS da je izdavalac ponovno poslao poruku.
type SubseqDelivType string

// CashDepositOperationType represents type of cash deposit
type CashDepositOperationType string

// TypeOfInv - represents the type of the invoice in question.
// TypeOfInv - vrsta stavke predstavlja vrstu stavke računa, npr. redovna prodaja ili vraćeni predmet.
type TypeOfInv string

// TypeOfSelfiss - invoice type issued by the customer.
// This element shows the type of the invoice that is issued by the customer himself.
type TypeOfSelfiss string

// InvNum - Invoice Number
type InvNum string

// InvOrdNum - Invoice Order Number
type InvOrdNum int64

// CorrectiveInvType - type of the corrective invoice.
// CorrectiveInvType - vrsta korektivnog računa.
type CorrectiveInvType string

// PayMethodType - payment method type
type PayMethodType string

// CurrencyCodeType - currency codes from ISO 4217 standard
type CurrencyCodeType string

// IDType represents Buyer/Seller identification type
type IDType string

// ExemptFromVATType - exemption from paying the VAT.
// ExemptFromVATType - izuzeće od PDV-a
type ExemptFromVATType string

// FeeType - Fee type.
type FeeType string

// Amount represents amount in given currency
type Amount float64

// CorrectiveInv - XML element groups data for an original invoice that will be corrected with current invoice.
// CorrectiveInv - XML elementi grupišu podatke za originalni račun koji će biti ispravljen trenutnim računom.
// IICRef - IKOF reference on the original invoice.
// IICRef - IKOF referenca na originalni račun
// IssueDateTime - Date and time the original invoice is created and issued at ENU.
// IssueDateTime - Datum i vrijeme kada je originalni račun kreiran i izdat od strane ENU.
// Type - Type of the corrective invoice.
// Type - Vrsta korektivnog računa
type CorrectiveInv struct {
	IICRef        string            `xml:"IICRef,attr"`
	IssueDateTime time.Time         `xml:"IssueDateTime,attr"`
	Type          CorrectiveInvType `xml:"Type,attr"`
}

// SumInvIICRef - XML element that contains one IKOF reference, e.g. reference of the invoice that is part of the summary invoice.
// SumInvIICRef - XML element koji sadrži jedan referentni IKOF, npr. referencu računa koji je dio zbirnog računa.
// IIC - IKOF of the invoice that is referenced in the summary invoice.
// IIC - IKOF računa na koji se poziva u zbirnom računu
// IssueDateTime - Date and time the invoice referenced by the summary invoice is created and issued at ENU.
// IssueDateTime - Datum i vrijeme kada je račun na koji se poziva u zbirnom računu kreiran i izdat od strane ENU.
type SumInvIICRef struct {
	IIC           string    `xml:"IIC,attr"`
	IssueDateTime time.Time `xml:"IssueDateTime,attr"`
}

// SupplyDateOrPeriod - XML element representing supply date or period of supply, if it is different from the date when the invoice was issued.
// SupplyDateOrPeriod - XML element koji predstavlja datum isporuke ili period isporuke, ako se razlikuje od datuma izdavanja računa.
// Start - Start day of the supply.
// Start - Datum početka isporuke
// End - End day of the supply.
// End - Posljednji dan isporuke
type SupplyDateOrPeriod struct {
	Start Date `xml:"Start,attr"`
	End   Date `xml:"End,attr"`
}

// Voucher - XML element representing single voucher.
// Voucher - Vaučer
// Num - Voucher serial number.
// Num - Serijski broj vaučera.
type Voucher struct {
	Num string `xml:"Num,attr"`
}

// VoucherNumber - XML element representing serial numbers of voucher sold.
// VoucherNumber - Lista serijskih brojeva vaučera.
// Voucher - XML element representing single voucher.
// Voucher - Vaučer
type VoucherNumber struct {
	Voucher Voucher `xml:"Voucher"`
}

// VoucherData - XML element representing data of vouchers sold
// VoucherData - Generalni podaci o prodanim vaučerima.
// D - (Date) Expiration date of the voucher.
// D - (Date) Datum važenja vaučera.
// N - (Nominal value) Nominal voucher value.
// N - (Nominal value) Nominalna vrijednost vaučera.
type VoucherData struct {
	D time.Time `xml:"D,attr"`
	N Amount    `xml:"N,attr"`
}

// PayMethod - XML element representing one payment method.
// PayMethod - XML element koji predstavlja jedan način plaćanja.
// Type - Type of the payment method.
// Type - Vrsta načina plaćanja.
// Amt - Amount paid by payment method in the currency in which the invoice was issued. It is mandatory if there are multiple payment methods.
// Amt - Iznos plaćen načinom plaćanja u valuti u kojoj je račun izdat. Obavezan je ako postoji više načina plaćanja.
// CompCard - Company card number if the payment method is company card. It is mandatory if the method / type of payment is stated: company.
// CompCard - Broj kartice preduzeća, ako je način plaćanja kartica preduzeća. Obavezno je ako je za način/vrstu plaćanja navedeno: preduzeće
// Vouchers - XML element that contains list of voucher numbers if the payment method is voucher. It is mandatory if the method / type of payment is stated: S voucher, M voucher.
// Vouchers - XML element koji sadrži listu brojeva vaučera, ako je način plaćanja vaučer. Obavezno je ako je za način/vrstu plaćanja navedeno: S VAUČER, M VAUČER
type PayMethod struct {
	Type     PayMethodType `xml:"Type,attr"`
	Amt      Amount        `xml:"Amt,attr"`
	CompCard string        `xml:"CompCard,attr,omitempty"`
	Vouchers *[]Voucher    `xml:"Vouchers>Voucher,omitempty"`
}

// Currency - XML element representing currency in which the amount on the invoice should be paid.
// Currency - XML element koji predstavlja valutu u kojoj je iznos na računu iskazan.
// Code - Currency code in which the amount on the invoice should be paid.
// Code - Kôd valute u kojoj je iznos na računu iskazan.
// ExRate - Exchange rate applied to calculate the equivalent amount of foreign currency for the total amount expressed in €. Exchange rate express equivalent amount of € for 1 unit of foreign currency.
// ExRate - Kurs razmjene primijenjen za izračunavanje ekvivalentnog iznosa ukupnog iznosa izraženog u valuti različitoj od €.
type Currency struct {
	Code   CurrencyCodeType `xml:"Code,attr"`
	ExRate float64          `xml:"ExRate,attr"`
}

// Seller - XML element representing seller’s data.
// Seller - XML element koji predstavlja podatke prodavca.
// IDType - Seller’s identification number type.
// IDType - Vrsta identifikacionog broja prodavca
// IDNum - Seller’s identification number.
// IDNum - Identifikacioni broj prodavca
// Name - Seller’s name.
// Name - Ime prodavca
// Address - Seller’s address.
// Address - Adresa prodavca. Obavezno ako je prodavac  stranac
// Town - Seller’s town.
// Town - Grad prodavca. Obavezno ako je prodavac  stranac.
// Country - Seller’s country.
// Country - Država prodavca. Obavezno ako je prodavac  stranac.
type Seller struct {
	IDType  IDType `xml:"IDType,attr"`
	IDNum   string `xml:"IDNum,attr"`
	Name    string `xml:"Name,attr,omitempty"`
	Address string `xml:"Address,attr,omitempty"`
	Town    string `xml:"Town,attr,omitempty"`
	Country string `xml:"Country,attr,omitempty"`
}

// Buyer - XML element representing buyer’s data.
// Buyer - XML element koji predstavlja podatke kupca.
// IDType - Buyer’s identification number type.
// IDType - Vrsta identifikacionog broja kupca. Obavezno ako postoji račun, vrsta samonaplatnog uređaja (automata). Obavezno ako postoji račun za iznos izvezene robe.
// IDNum - Buyer’s identification number.
// IDNum - Identifikacioni broj kupca. Obavezno ako postoji račun, vrsta samonaplatnog uređaja (automata). Obavezno ako postoji račun za iznos izvezene robe.
// Name - Buyer’s name.
// Name - Ime kupca. Obavezno ako postoji račun, kupac, vrsta identifikacije i račun, kupac identifikacioni broj
// Address - Buyer’s address.
// Address - Adresa kupca. Obavezno ako postoji račun, kupac, vrsta identifikacije i račun, kupac identifikacioni broj.
// Town - Buyer’s town.
// Town - Grad kupca. Obavezno ako postoji račun, kupac, vrsta identifikacije i račun, kupac identifikacioni broj.
// Country - Buyer’s country.
// Country - Država prodavca. Obavezno ako postoji račun, kupac, vrsta identifikacije i račun, kupac identifikacioni broj.
type Buyer struct {
	IDType  IDType `xml:"IDType,attr"`
	IDNum   string `xml:"IDNum,attr"`
	Name    string `xml:"Name,attr,omitempty"`
	Address string `xml:"Address,attr,omitempty"`
	Town    string `xml:"Town,attr,omitempty"`
	Country string `xml:"Country,attr,omitempty"`
}

// Fee - XML element representing one fee.
// Fee - XML element koji predstavlja jednu naknadu
// Type - Type of the fee.
// Type - Vrsta naknade
// Amt - Amount of the fee.
// Amt - Iznos naknade
type Fee struct {
	Type FeeType `xml:"Type,attr"`
	Amt  Amount  `xml:"Amt,attr"`
}

// SameTax - XML element representing one same tax item.
// SameTax - XML element koji predstavlja jednu istu poresku stavku.
// NumOfItems - Number of items.
// NumOfItems - Broj stavki
// PriceBefVAT - Price before VAT.
// PriceBefVAT - Cijena bez PDV-a
// VATRate - VAT rate.
// VATRate - Stopa PDV-a
// ExemptFromVAT - Exempt from VAT.
// ExemptFromVAT - Izuzeto od PDV-a
// VATAmt - VAT amount.
// VATAmt - Iznos PDV-a
// Fees - XML element representing list of fees.
// Fees - XML element koji predstavlja listu naknada
type SameTax struct {
	NumOfItems    int64             `xml:"NumOfItems,attr"`
	PriceBefVAT   Amount            `xml:"PriceBefVAT,attr"`
	VATRate       Amount            `xml:"VATRate,attr"`
	ExemptFromVAT ExemptFromVATType `xml:"ExemptFromVAT,attr,omitempty"`
	VATAmt        Amount            `xml:"VATAmt,attr"`
}

// BadDebtInv - XML element groups data for an original invoice that will be declared bad debt invoice, as uncollectible.
// BadDebtInv - XML elementi grupišu podatke za originalni račun koji će biti proglašen nenaplativim s trenutnim računom.
// IICRef - IKOF reference on the original invoice.
// IICRef - IKOF referenca na originalni račun
// IssueDateTime - Date and time the original invoice is created and issued at ENU.
// IssueDateTime - Datum i vrijeme kreiranja i izdavanja originalnog računa od strane ENU.
type BadDebtInv struct {
	IICRef        string    `xml:"IICRef,attr"`
	IssueDateTime time.Time `xml:"IssueDateTime,attr"`
}

// Item - XML element representing one item.
// Item - XML element koji predstavlja jednu stavku.
// N - (Name) Name of the item (goods or services).
// N - (Name) Naziv stavke (roba ili usluge)
// C - (Code) Code of the item from the barcode or similar representation.
// C - (Code) Kôd stavke iz bar koda ili slično prikazivanje.
// IN - (Is Investment) If the item is investment for the buyer.
// IN - (Is Investment) Ako je stavka investicija.
// U - (Unit of measure) What is the item’s unit of measure (piece, weight measure, length measure, etc.)
// U - (Unit of measure) Koja je jedinica mjere stavke (komad, jedinica za mjerenje težine, jedinica za mjerenje dužine, itd.)
// VS - (Voucher sold) XML element representing vouchers sold
// VS - (Voucher sold) Lista prodanih vaučera na stavki.
// VN - (Voucher sold numbers) XML element representing serial numbers of voucher sold.
// VN - (Voucher sold numbers) Lista serijskih brojeva vaučera.
// Q - (Quantity) Amount or number (quantity) of items.
// Q - (Quantity) Količina ili broj stavki.
// UPB - (Unit price without VAT) Unit price before VAT is applied.
// UPB - (Unit price without VAT) Jedinična cijena prije dodavanja PDV-a
// UPA - (Unique price with VAT) Unit price after VAT is applied.
// UPA - (Unique price with VAT) Jedinična cijena poslije dodavanja PDV-a
// R - (Rebate) Percentage of the rebate.
// R - (Rebate) Procenat rabata.
// RR - (Rebate reducing base price) Is rebate reducing tax base amount?
// RR - (Rebate reducing base price) Da li rabat smanjuje osnovni iznos?
// PB - (Price before VAT) Total price of goods and services before the tax.
// PB - (Price before VAT) Ukupna cijena robe i usluga prije oporezivanja
// VR - (VAT rate) Rate of VAT. Mandatory if issuer is in PDV system.
// VR - (VAT rate) Stopa PDV-a. Obavezno ako je izdavalac u sistemu PDV-a.
// EX - (Excempt from VAT) Exempt from VAT.
// EX - (Exempt from VAT) Izuzeće od plaćanja PDV-a.
// VA - (VAT amount) Amount of VAT for goods and services. Mandatory if the issuer is in the VAT system. Mandatory if there is a self-charging device (vending machine) and the issuer is in the VAT system. Mandatory if reverse charging applies.
// VA - (VAT amount) Iznos PDV-a za robu i usluge. Obavezno ako je izdavalac u sistemu PDV-a.
// PA - (Price after applying VAT) Total price of goods after the tax and applying discounts.
// PA - (Price after applying VAT) Ukupna cijena robe nakon oporezivanja i primjene popusta.
type Item struct {
	N   string            `xml:"N,attr"`
	C   string            `xml:"C,attr,omitempty"`
	IN  bool              `xml:"IN,attr,omitempty"`
	U   string            `xml:"U,attr"`
	VS  *[]VoucherData    `xml:"VS>VD,omitempty"`
	VN  *[]Voucher        `xml:"VN>V,omitempty"`
	Q   Amount            `xml:"Q,attr"`
	UPB Amount            `xml:"UPB,attr"`
	UPA Amount            `xml:"UPA,attr"`
	R   Amount            `xml:"R,attr,omitempty"`
	RR  bool              `xml:"RR,attr,omitempty"`
	PB  Amount            `xml:"PB,attr"`
	VR  Amount            `xml:"VR,attr"`
	EX  ExemptFromVATType `xml:"EX,attr,omitempty"`
	VA  Amount            `xml:"VA,attr"`
	PA  Amount            `xml:"PA,attr"`
}

// MarshalXMLAttr allows to generate "xxx.xx" format
func (a Amount) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {

	return xml.Attr{
		Name:  name,
		Value: fmt.Sprintf("%.02f", a),
	}, nil
}

// Invoice represents details of invoice
type Invoice struct {
	InvType            InvType             `xml:"InvType,attr"`
	TypeOfInv          TypeOfInv           `xml:"TypeOfInv,attr"`
	TypeOfSelfiss      TypeOfSelfiss       `xml:"TypeOfSelfiss,attr,omitempty"`
	IssueDateTime      DateTime            `xml:"IssueDateTime,attr"`
	InvNum             InvNum              `xml:"InvNum,attr"`
	InvOrdNum          InvOrdNum           `xml:"InvOrdNum,attr"`
	TCRCode            TCRCode             `xml:"TCRCode,attr"`
	IsIssuerInVAT      bool                `xml:"IsIssuerInVAT,attr"`
	TaxFreeAmt         Amount              `xml:"TaxFreeAmt,attr,omitempty"`
	MarkupAmt          Amount              `xml:"MarkupAmt,attr,omitempty"`
	GoodsExAmt         Amount              `xml:"GoodsExAmt,attr,omitempty"`
	TotPriceWoVAT      Amount              `xml:"TotPriceWoVAT,attr"`
	TotVATAmt          Amount              `xml:"TotVATAmt,attr"`
	TotPrice           Amount              `xml:"TotPrice,attr"`
	OperatorCode       OperatorCode        `xml:"OperatorCode,attr"`
	BusinUnitCode      BusinUnitCode       `xml:"BusinUnitCode,attr"`
	SoftCode           SoftCode            `xml:"SoftCode,attr"`
	IIC                string              `xml:"IIC,attr"`
	IICSignature       string              `xml:"IICSignature,attr"`
	IsReverseCharge    bool                `xml:"IsReverseCharge,attr"`
	PayDeadline        *DateTime           `xml:"PayDeadline,attr,omitempty"`
	ParagonBlockNum    string              `xml:"ParagonBlockNum,attr,omitempty"`
	CorrectiveInv      *CorrectiveInv      `xml:"CorrectiveInv,omitempty"`
	SumInvIICRefs      *[]SumInvIICRef     `xml:"SumInvIICRefs>SumInvIICRef,omitempty"`
	SupplyDateOrPeriod *SupplyDateOrPeriod `xml:"SupplyDateOrPeriod,omitempty"`
	PayMethods         []*PayMethod        `xml:"PayMethods>PayMethod"`
	Currency           *Currency           `xml:"Currency,omitempty"`
	Seller             *Seller             `xml:"Seller"`
	Buyer              *Buyer              `xml:"Buyer"`
	Items              *[]*Item            `xml:"Items>I"`
	SameTaxes          *[]*SameTax         `xml:"SameTaxes>SameTax"`
	Fees               *[]Fee              `xml:"Fees>Fee,omitempty"`
	BadDebtInv         *BadDebtInv         `xml:"BadDebtInv,omitempty"`
}

// Header represents common struct of header of any request
type Header struct {
	UUID            string          `xml:"UUID,attr"`
	SendDateTime    DateTime        `xml:"SendDateTime,attr"`
	SubseqDelivType SubseqDelivType `xml:"SubseqDelivType,attr,omitempty"`
}

// Fault represents common structure of faulted response details
type Fault struct {
	FaultCode   string `xml:"faultcode,omitempty"`
	FaultString string `xml:"faultstring,omitempty"`
	Detail      struct {
		Code         int    `xml:"code,omitempty"`
		RequestUUID  string `xml:"requestUUID,omitempty"`
		ResponseUUID string `xml:"responseUUID,omitempty"`
	}
}

// CanonicalizationMethod represents details about CanonicalizationMethod
type CanonicalizationMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// SignatureMethod represents details about SignatureMethod
type SignatureMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// Transform represents details about Transform
type Transform struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// DigestMethod represents details about DigestMethod
type DigestMethod struct {
	Algorithm string `xml:"Algorithm,attr"`
}

// Reference represents details about Reference
type Reference struct {
	URI          string       `xml:"URI,attr"`
	Transforms   []Transform  `xml:"Transforms>Transform"`
	DigestMethod DigestMethod `xml:"DigestMethod"`
	DigestValue  string       `xml:"DigestValue"`
}

// X509Data represents details about X509Data
type X509Data struct {
	X509Certificate string `xml:"X509Certificate"`
	X509SubjectName string `xml:"X509SubjectName"`
}

// KeyInfo represents details about KeyInfo
type KeyInfo struct {
	X509Data X509Data `xml:"X509Data"`
}

// SignedInfo represents details about SignedInfo
type SignedInfo struct {
	CanonicalizationMethod CanonicalizationMethod `xml:"CanonicalizationMethod"`
	SignatureMethod        SignatureMethod        `xml:"SignatureMethod"`
	Reference              Reference              `xml:"Reference"`
}

// Signature represents common signature structure
type Signature struct {
	XMLName        xml.Name   `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
	SignedInfo     SignedInfo `xml:"SignedInfo"`
	SignatureValue string     `xml:"SignatureValue"`
	KeyInfo        KeyInfo    `xml:"KeyInfo"`
}
