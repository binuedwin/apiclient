package pkg

type CreditPartyIdentifierRequestWrapper struct {
	CreditPartyIdentifier CreditPartyIdentifier `json:"credit_party_identifier"`
}

type CreditPartyIdentifier struct {
	MSISDN            string `json:"msisdn,omitempty"`
	BankAccountNumber string `json:"bank_account_number,omitempty"`
	IBAN              string `json:"iban,omitempty"`
	CLABE             string `json:"clabe,omitempty"`
	CBU               string `json:"cbu,omitempty"`
	CBUALIAS          string `json:"cbu_alias,omitempty"`
	SwiftBICCode      string `json:"swift_bic_code,omitempty"`
	BIKCode           string `json:"bik_code,omitempty"`
	IFSCode           string `json:"ifs_code,omitempty"`
	SortCode          string `json:"sort_code,omitempty"`
	ABARoutingNumber  string `json:"aba_routing_number,omitempty"`
	BSBNumber         string `json:"bsb_number,omitempty"`
	BranchNumber      string `json:"branch_cumber,omitempty"`
	RoutingCode       string `json:"routing_code,omitempty"`
	EntityTTID        int    `json:"entity_tt_id,omitempty"`
	AccountType       string `json:"account_type,omitempty"`
	AccountNumber     string `json:"account_number,omitempty"`
	Email             string `json:"email,omitempty"`
}

// ReceivingBusinessInformation represents receiving business information for a given transaction of type C2B or B2B.
type ReceivingBusinessInformation struct {
	RegisteredName                 *string `json:"registered_name"`
	TradingName                    *string `json:"trading_name"`
	Address                        *string `json:"address"`
	PostalCode                     *string `json:"postal_code"`
	City                           *string `json:"city"`
	ProvinceState                  *string `json:"province_state"`
	CountryIsoCode                 *string `json:"country_iso_code"` // format: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
	Msisdn                         *string `json:"msisdn"`
	Email                          *string `json:"email"`
	RegistrationNumber             *string `json:"registration_number"`
	TaxID                          *string `json:"tax_id"`
	DateOfIncorporation            *string `json:"date_of_incorporation"` // format: https://en.wikipedia.org/wiki/ISO_8601
	RepresentativeLastname         *string `json:"representative_lastname"`
	RepresentativeLastname2        *string `json:"representative_lastname2"`
	RepresentativeFirstname        *string `json:"representative_firstname"`
	RepresentativeMiddlename       *string `json:"representative_middlename"`
	RepresentativeNativename       *string `json:"representative_nativename"`
	RepresentativeIDType           *string `json:"representative_id_type"`
	RepresentativeIDCountryIsoCode *string `json:"representative_id_country_iso_code"` // format: https://en.wikipedia.org/wiki/ISO_8601
	RepresentativeIDNumber         *string `json:"representative_id_number"`
	RepresentativeIDDeliveryDate   *string `json:"representative_id_delivery_date"`   // format: https://en.wikipedia.org/wiki/ISO_8601
	RepresentativeIDExpirationDate *string `json:"representative_id_expiration_date"` // format: https://en.wikipedia.org/wiki/ISO_8601
}

// Beneficiary represents beneficiary information for a given transaction of type C2C or B2C.
type Beneficiary struct {
	LastName              *string `json:"last_name"`
	LastName2             *string `json:"last_name2"`
	MiddleName            *string `json:"middle_name"`
	FirstName             *string `json:"first_name"`
	NativeName            *string `json:"native_name"`
	NationalityCountryI   *string `json:"nationality_country_i"` // format: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
	Code                  *string `json:"code"`
	DateOfBirth           *string `json:"date_of_birth"`             // format: https://en.wikipedia.org/wiki/ISO_8601
	CountryOfBirthISOCode *string `json:"country_of_birth_iso_code"` // format: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
	Gender                *string `json:"gender"`                    // enum: 'MALE', 'FEMALE'
	Address               *string `json:"address"`
	PostalCode            *string `json:"postal_code"`
	City                  *string `json:"city"`
	CountryISOCode        *string `json:"country_iso_code"`
	MSISDN                *string `json:"msisdn"`
	Email                 *string `json:"email"`
	IDType                *string `json:"id_type"`
	IDCountryISOCode      *string `json:"id_country_iso_code"` // format: https://en.wikipedia.org/wiki/ISO_3166-1_alpha-3
	IDNumber              *string `json:"id_number"`
	IDDeliveryDate        *string `json:"id_delivery_date"`   // format: https://en.wikipedia.org/wiki/ISO_8601
	IDExpirationDate      *string `json:"id_expiration_date"` // format: https://en.wikipedia.org/wiki/ISO_8601
	Occupation            *string `json:"occupation"`
	BankAccountHolderName *string `json:"bank_account_holder_name"`
	ProvinceState         *string `json:"province_state"`
}
