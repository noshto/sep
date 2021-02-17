package sep

// Service URLs
const (
	TestingURL          = "https://efitest.tax.gov.me/fs-v1/FiscalizationService.wsdl"
	TestingVerifyURL    = "https://efitest.tax.gov.me/ic/#/verify?iic=%s&tin=%s&crtd=%s&ord=%d&bu=%s&cr=%s&sw=%s&prc=%s"
	ProductionURL       = "https://efi.tax.gov.me/fs-v1/FiscalizationService.wsdl"
	ProductionVerifyURL = "https://mapr.tax.gov.me/ic/#/verify?iic=%s&tin=%s&crtd=%s&ord=%d&bu=%s&cr=%s&sw=%s&prc=%s"
)

// OTHER represents default "OTHER" option that can be used for ExemptFromVATType, FeeType, PayMethodType and TypeOfSelfiss
const (
	OTHER = "OTHER"
)

// Possible values for CashDepositOperationType
const (
	INITIAL  CashDepositOperationType = "INITIAL"
	WITHDRAW CashDepositOperationType = "WITHDRAW"
)

// Possible values for CorrectiveInvType
const (
	CORRECTIVE CorrectiveInvType = "CORRECTIVE"
	DEBIT      CorrectiveInvType = "DEBIT"
	CREDIT     CorrectiveInvType = "CREDIT"
)

// Possible values for CurrencyType
const (
	// United Arab Emirates Dirham
	AED CurrencyCodeType = "AED"

	// Afghanistan Afghani
	AFN CurrencyCodeType = "AFN"

	// Armenia Dram
	AMD CurrencyCodeType = "AMD"

	// Netherlands Antilles Guilder
	ANG CurrencyCodeType = "ANG"

	// Angola Kwanza
	AOA CurrencyCodeType = "AOA"

	// Argentina Peso
	ARS CurrencyCodeType = "ARS"

	// Australia Dollar
	AUD CurrencyCodeType = "AUD"

	// Aruba Guilder
	AWG CurrencyCodeType = "AWG"

	// Azerbaijan Manat
	AZN CurrencyCodeType = "AZN"

	// Bosnia and Herzegovina Convertible Mark
	BAM CurrencyCodeType = "BAM"

	// Barbados Dollar
	BBD CurrencyCodeType = "BBD"

	// Bangladesh Taka
	BDT CurrencyCodeType = "BDT"

	// Bulgaria Lev
	BGN CurrencyCodeType = "BGN"

	// Bahrain Dinar
	BHD CurrencyCodeType = "BHD"

	// Burundi Franc
	BIF CurrencyCodeType = "BIF"

	// Bermuda Dollar
	BMD CurrencyCodeType = "BMD"

	// Brunei Darussalam Dollar
	BND CurrencyCodeType = "BND"

	// Bolivia Boliviano
	BOB CurrencyCodeType = "BOB"

	// Brazil Real
	BRL CurrencyCodeType = "BRL"

	// Bahamas Dollar
	BSD CurrencyCodeType = "BSD"

	// Bhutan Ngultrum
	BTN CurrencyCodeType = "BTN"

	// Botswana Pula
	BWP CurrencyCodeType = "BWP"

	// Belarus Ruble
	BYN CurrencyCodeType = "BYN"

	// Belize Dollar
	BZD CurrencyCodeType = "BZD"

	// Canada Dollar
	CAD CurrencyCodeType = "CAD"

	// Congo/Kinshasa Franc
	CDF CurrencyCodeType = "CDF"

	// Switzerland Franc
	CHF CurrencyCodeType = "CHF"

	// Chile Peso
	CLP CurrencyCodeType = "CLP"

	// China Yuan Renminbi
	CNY CurrencyCodeType = "CNY"

	// Colombia Peso
	COP CurrencyCodeType = "COP"

	// Costa Rica Colon
	CRC CurrencyCodeType = "CRC"

	// Cuba Convertible Peso
	CUC CurrencyCodeType = "CUC"

	// Cuba Peso
	CUP CurrencyCodeType = "CUP"

	// Cape Verde Escudo
	CVE CurrencyCodeType = "CVE"

	// Czech Republic Koruna
	CZK CurrencyCodeType = "CZK"

	// Djibouti Franc
	DJF CurrencyCodeType = "DJF"

	// Denmark Krone
	DKK CurrencyCodeType = "DKK"

	// Dominican Republic Peso
	DOP CurrencyCodeType = "DOP"

	// Algeria Dinar
	DZD CurrencyCodeType = "DZD"

	// Egypt Pound
	EGP CurrencyCodeType = "EGP"

	// Eritrea Nakfa
	ERN CurrencyCodeType = "ERN"

	// Ethiopia Birr
	ETB CurrencyCodeType = "ETB"

	// Euro Member Countries
	EUR CurrencyCodeType = "EUR"

	// Fiji Dollar
	FJD CurrencyCodeType = "FJD"

	// Falkland Islands (Malvinas) Pound
	FKP CurrencyCodeType = "FKP"

	// United Kingdom Pound
	GBP CurrencyCodeType = "GBP"

	// Georgia Lari
	GEL CurrencyCodeType = "GEL"

	// Guernsey Pound
	GGP CurrencyCodeType = "GGP"

	// Ghana Cedi
	GHS CurrencyCodeType = "GHS"

	// Gibraltar Pound
	GIP CurrencyCodeType = "GIP"

	// Gambia Dalasi
	GMD CurrencyCodeType = "GMD"

	// Guinea Franc
	GNF CurrencyCodeType = "GNF"

	// Guatemala Quetzal
	GTQ CurrencyCodeType = "GTQ"

	// Guyana Dollar
	GYD CurrencyCodeType = "GYD"

	// Hong Kong Dollar
	HKD CurrencyCodeType = "HKD"

	// Honduras Lempira
	HNL CurrencyCodeType = "HNL"

	// Croatia Kuna
	HRK CurrencyCodeType = "HRK"

	// Haiti Gourde
	HTG CurrencyCodeType = "HTG"

	// Hungary Forint
	HUF CurrencyCodeType = "HUF"

	// Indonesia Rupiah
	IDR CurrencyCodeType = "IDR"

	// Israel Shekel
	ILS CurrencyCodeType = "ILS"

	// Isle of Man Pound
	IMP CurrencyCodeType = "IMP"

	// India Rupee
	INR CurrencyCodeType = "INR"

	// Iraq Dinar
	IQD CurrencyCodeType = "IQD"

	// Iran Rial
	IRR CurrencyCodeType = "IRR"

	// Iceland Krona
	ISK CurrencyCodeType = "ISK"

	// Jersey Pound
	JEP CurrencyCodeType = "JEP"

	// Jamaica Dollar
	JMD CurrencyCodeType = "JMD"

	// Jordan Dinar
	JOD CurrencyCodeType = "JOD"

	// Japan Yen
	JPY CurrencyCodeType = "JPY"

	// Kenya Shilling
	KES CurrencyCodeType = "KES"

	// Kyrgyzstan Som
	KGS CurrencyCodeType = "KGS"

	// Cambodia Riel
	KHR CurrencyCodeType = "KHR"

	// Comorian Franc
	KMF CurrencyCodeType = "KMF"

	// Korea (North) Won
	KPW CurrencyCodeType = "KPW"

	// Korea (South) Won
	KRW CurrencyCodeType = "KRW"

	// Kuwait Dinar
	KWD CurrencyCodeType = "KWD"

	// Cayman Islands Dollar
	KYD CurrencyCodeType = "KYD"

	// Kazakhstan Tenge
	KZT CurrencyCodeType = "KZT"

	// Laos Kip
	LAK CurrencyCodeType = "LAK"

	// Lebanon Pound
	LBP CurrencyCodeType = "LBP"

	// Sri Lanka Rupee
	LKR CurrencyCodeType = "LKR"

	// Liberia Dollar
	LRD CurrencyCodeType = "LRD"

	// Lesotho Loti
	LSL CurrencyCodeType = "LSL"

	// Libya Dinar
	LYD CurrencyCodeType = "LYD"

	// Morocco Dirham
	MAD CurrencyCodeType = "MAD"

	// Moldova Leu
	MDL CurrencyCodeType = "MDL"

	// Madagascar Ariary
	MGA CurrencyCodeType = "MGA"

	// Macedonia Denar
	MKD CurrencyCodeType = "MKD"

	// Myanmar (Burma) Kyat
	MMK CurrencyCodeType = "MMK"

	// Mongolia Tughrik
	MNT CurrencyCodeType = "MNT"

	// Macau Pataca
	MOP CurrencyCodeType = "MOP"

	// Mauritania Ouguiya
	MRU CurrencyCodeType = "MRU"

	// Mauritius Rupee
	MUR CurrencyCodeType = "MUR"

	// Maldives (Maldive Islands) Rufiyaa
	MVR CurrencyCodeType = "MVR"

	// Malawi Kwacha
	MWK CurrencyCodeType = "MWK"

	// Mexico Peso
	MXN CurrencyCodeType = "MXN"

	// Malaysia Ringgit
	MYR CurrencyCodeType = "MYR"

	// Mozambique Metical
	MZN CurrencyCodeType = "MZN"

	// Namibia Dollar
	NAD CurrencyCodeType = "NAD"

	// Nigeria Naira
	NGN CurrencyCodeType = "NGN"

	// Nicaragua Cordoba
	NIO CurrencyCodeType = "NIO"

	// Norway Krone
	NOK CurrencyCodeType = "NOK"

	// Nepal Rupee
	NPR CurrencyCodeType = "NPR"

	// New Zealand Dollar
	NZD CurrencyCodeType = "NZD"

	// Oman Rial
	OMR CurrencyCodeType = "OMR"

	// Panama Balboa
	PAB CurrencyCodeType = "PAB"

	// Peru Sol
	PEN CurrencyCodeType = "PEN"

	// Papua New Guinea Kina
	PGK CurrencyCodeType = "PGK"

	// Philippines Peso
	PHP CurrencyCodeType = "PHP"

	// Pakistan Rupee
	PKR CurrencyCodeType = "PKR"

	// Poland Zloty
	PLN CurrencyCodeType = "PLN"

	// Paraguay Guarani
	PYG CurrencyCodeType = "PYG"

	// Qatar Riyal
	QAR CurrencyCodeType = "QAR"

	// Romania Leu
	RON CurrencyCodeType = "RON"

	// Serbia Dinar
	RSD CurrencyCodeType = "RSD"

	// Russia Ruble
	RUB CurrencyCodeType = "RUB"

	// Rwanda Franc
	RWF CurrencyCodeType = "RWF"

	// Saudi Arabia Riyal
	SAR CurrencyCodeType = "SAR"

	// Solomon Islands Dollar
	SBD CurrencyCodeType = "SBD"

	// Seychelles Rupee
	SCR CurrencyCodeType = "SCR"

	// Sudan Pound
	SDG CurrencyCodeType = "SDG"

	// Sweden Krona
	SEK CurrencyCodeType = "SEK"

	// Singapore Dollar
	SGD CurrencyCodeType = "SGD"

	// Saint Helena Pound
	SHP CurrencyCodeType = "SHP"

	// Sierra Leone Leone
	SLL CurrencyCodeType = "SLL"

	// Somalia Shilling
	SOS CurrencyCodeType = "SOS"

	// Seborga Luigino
	SPL CurrencyCodeType = "SPL"

	// Suriname Dollar
	SRD CurrencyCodeType = "SRD"

	// Sao Tome and Principe Dobra
	STN CurrencyCodeType = "STN"

	// El Salvador Colon
	SVC CurrencyCodeType = "SVC"

	// Syria Pound
	SYP CurrencyCodeType = "SYP"

	// eSwatini Lilangeni
	SZL CurrencyCodeType = "SZL"

	// Thailand Baht
	THB CurrencyCodeType = "THB"

	// Tajikistan Somoni
	TJS CurrencyCodeType = "TJS"

	// Turkmenistan Manat
	TMT CurrencyCodeType = "TMT"

	// Tunisia Dinar
	TND CurrencyCodeType = "TND"

	// Tonga Pa'anga
	TOP CurrencyCodeType = "TOP"

	// Turkey Lira
	TRY CurrencyCodeType = "TRY"

	// Trinidad and Tobago Dollar
	TTD CurrencyCodeType = "TTD"

	// Tuvalu Dollar
	TVD CurrencyCodeType = "TVD"

	// Taiwan New Dollar
	TWD CurrencyCodeType = "TWD"

	// Tanzania Shilling
	TZS CurrencyCodeType = "TZS"

	// Ukraine Hryvnia
	UAH CurrencyCodeType = "UAH"

	// Uganda Shilling
	UGX CurrencyCodeType = "UGX"

	// United States Dollar
	USD CurrencyCodeType = "USD"

	// Uruguay Peso
	UYU CurrencyCodeType = "UYU"

	// Uzbekistan Som
	UZS CurrencyCodeType = "UZS"

	// Venezuela Bolivar
	VEF CurrencyCodeType = "VEF"

	// Viet Nam Dong
	VND CurrencyCodeType = "VND"

	// Vanuatu Vatu
	VUV CurrencyCodeType = "VUV"

	// Samoa Tala
	WST CurrencyCodeType = "WST"

	// Communaute Financiere Africaine (BEAC) CFA Franc BEAC
	XAF CurrencyCodeType = "XAF"

	// East Caribbean Dollar
	XCD CurrencyCodeType = "XCD"

	// International Monetary Fund (IMF) Special Drawing Rights
	XDR CurrencyCodeType = "XDR"

	// Communaute Financiere Africaine (BCEAO) Franc
	XOF CurrencyCodeType = "XOF"

	// Comptoirs Francais du Pacifique (CFP) Franc
	XPF CurrencyCodeType = "XPF"

	// Yemen Rial
	YER CurrencyCodeType = "YER"

	// South Africa Rand
	ZAR CurrencyCodeType = "ZAR"

	// Zambia Kwacha
	ZMW CurrencyCodeType = "ZMW"

	// Zimbabwe Dollar
	ZWD CurrencyCodeType = "ZWD"
)

// Possible values of ExemptFromVATType
// VAT_CL17 - Mjesto izdavanja računa
// VAT_CL20 - Poreska osnovica i ispravka poreske osnovice
// VAT_CL26 - Disengagement from public interest
// VAT_CL26 - Oslobođenja od javnog interesa
// VAT_CL27 - Other disengagements
// VAT_CL27 - Ostala oslobođenja
// VAT_CL28 - Disengagement when importing goods
// VAT_CL28 - Oslobođenja kod uvoza proizvoda
// VAT_CL29 - Disengagement when importing goods temporarily
// VAT_CL29 - Oslobođenja kod privremenog uvoza proizvoda
// VAT_CL30 - Special disengagements
// VAT_CL30 - Posebna oslobođenja
const (
	CL17 ExemptFromVATType = "VAT_CL17"
	CL20 ExemptFromVATType = "VAT_CL20"
	CL26 ExemptFromVATType = "VAT_CL26"
	CL27 ExemptFromVATType = "VAT_CL27"
	CL28 ExemptFromVATType = "VAT_CL28"
	CL29 ExemptFromVATType = "VAT_CL29"
	CL30 ExemptFromVATType = "VAT_CL30"
)

// Possible values for FeeType
// FeeTypePACK - Packing fee
// FeeTypePACK - Naknada za pakovanje
// FeeTypeBOTTLE - Return glass bottles fee
// FeeTypeBOTTLE - Naknada za povrat staklenih flaša
// FeeTypeCOMMISION - Money exchange office commission
// FeeTypeCOMMISION - Provizija za poslove mjenjačnice
// FeeTypeOTHER - Other fees
// FeeTypeOTHER - Ostale naknade koje ovdje nisu navedene.
const (
	PACK      FeeType = "PACK"
	BOTTLE    FeeType = "BOTTLE"
	COMMISION FeeType = "COMMISION"
	// OTHER     FeeType = "OTHER"
)

// Possible values for the IDType
// IDTypeTIN - Personal tax  number
// IDTypeTIN - PIB/JMB broj
// IDTypeID - Personal identification number
// IDTypeID - Lični matični broj
// IDTypePASS - Passport number
// IDTypePASS - Broj pasoša
// IDTypeVAT - VAT number
// IDTypeVAT - PDV broj
// IDTypeTAX - Tax number
// IDTypeTAX - Poreski broj
// IDTypeSOC - Social security number
// IDTypeSOC - Broj socijalnog osiguranja
const (
	IDTypeTIN  IDType = "TIN"
	IDTypeID   IDType = "ID"
	IDTypePASS IDType = "PASS"
	IDTypeVAT  IDType = "VAT"
	IDTypeTAX  IDType = "TAX"
	IDTypeSOC  IDType = "SOC"
)

// Possible values for PayMethodType
// BANKNOTE - Notes and coins
// BANKNOTE - Novčanice i kovanice
// Allowed only for: CASH
//
// CARD - Credit or debit card
// CARD - Kreditna i debitna kartica
// Allowed only for: CASH
//
// BUSINESSCARD - Bank cheque
// BUSINESSCARD - Kreditna i debitna kartica banke izdata poreskom obvezniku
// Allowed only for: NONCASH
//
// SVOUCHER - Onetime voucher
// SVOUCHER - Jednokratni vaučer
// Allowed only for: NONCASH
//
// COMPANY - Company cards
// COMPANY - Kartice preduzeća prodavca i slično
// Allowed only for: CASH
//
// ORDER - Invoice to be paid in summary invoice
// ORDER - Račun još nije plaćen. Bite plaćen zbirnim računom
// Allowed only for: CASH and NONCASH
//
// ADVANCE - Advance payment
// ADVANCE - Plaćanje avansom
// Allowed only for: NONCASH
//
// ACCOUNT - Transaction account
// ACCOUNT - Transakcioni račun
// Allowed only for: NONCASH
//
// FACTORING - Factoring
// FACTORING - Faktoring
// Allowed only for: NONCASH
//
// OTHER - Ostala bezgotovinska plaćanja
// Allowed only for: NONCASH
//
// OTHER-CASH - Ostala gotovinska plaćanja
// Allowed only for: CASH
const (
	BANKNOTE           PayMethodType = "BANKNOTE"
	CARD               PayMethodType = "CARD"
	BUSINESSCARD       PayMethodType = "BUSINESSCARD"
	SVOUCHER           PayMethodType = "SVOUCHER"
	COMPANY            PayMethodType = "COMPANY"
	ORDER              PayMethodType = "ORDER"
	ADVANCE            PayMethodType = "ADVANCE"
	ACCOUNT            PayMethodType = "ACCOUNT"
	FACTORING          PayMethodType = "FACTORING"
	PayMethodTypeOTHER PayMethodType = "OTHER"
	OTHER_CASH         PayMethodType = "OTHER-CASH"
)

// Possible values for SubseqDelivType
// NOINTERNET - When TCR operates in the area where there is no Internet available.
// NOINTERNET - Ako ENU djeluje u području bez interneta
// BOUNDBOOK - When TCR is not working and message cannot be created with TCR.
// BOUNDBOOK - ENU ne radi i ne može se kreirati poruka
// SERVICE - When there is an issue with the fiscalization service that blocks fiscalization.
// SERVICE - Problem sa fiskalnim servisom
// TECHNICALERROR - When there is a temporary technical error at TCR side that prevents successful fiscalization.
// TECHNICALERROR - Tehnička greška
// BUSINESSNEEDS - When there is a subsequent sending conditioned by the way of doing business that prevents successful fiscalization.
// BUSINESSNEEDS - Naknadno slanje uslovljeno načinom poslovanja
const (
	NOINTERNET     SubseqDelivType = "NOINTERNET"
	BOUNDBOOK      SubseqDelivType = "BOUNDBOOK"
	SERVICE        SubseqDelivType = "SERVICE"
	TECHNICALERROR SubseqDelivType = "TECHNICALERROR"
	BUSINESSNEEDS  SubseqDelivType = "BUSINESSNEEDS"
)

// Possible values for TCRType
// REGULAR - Standard ENU
// VENDING - Self-vending machine
const (
	REGULAR TCRType = "REGULAR"
	VENDING TCRType = "VENDING"
)

// Possible values for TypeOfInv
// CASH - Cash payment
// Gotovinsko plaćanje
// NONCASH - Non-cash payment
// Bezgotovinsko plaćanje
const (
	CASH    TypeOfInv = "CASH"
	NONCASH TypeOfInv = "NONCASH"
)

// Possible values for TypeOfSelfiss
// AGREEMENT - Previous agreement by the sides
// AGREEMENT - Prethodni sporazum strana.
// DOMESTIC - Buying from local
// DOMESTIC - Otkup od domaćih poljoprivrednika.
// ABROAD - Buying services abroad.
// ABROAD - Nabavka usluga iz inostranstva.
// OTHER - Other
// OTHER - Ostalo
const (
	AGREEMENT TypeOfSelfiss = "AGREEMENT"
	DOMESTIC  TypeOfSelfiss = "DOMESTIC"
	ABROAD    TypeOfSelfiss = "ABROAD"
	// TypeOfSelfissOTHER     TypeOfSelfiss = "OTHER"
)
