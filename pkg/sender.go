package pkg

// Sender represents sender information for a given transaction of type C2C or C2B.
type Sender struct {
	LastName                  *string `json:"lastname"`
	LastName2                 *string `json:"lastname2"`
	MiddleName                *string `json:"middlename"`
	FirstName                 *string `json:"firstname"`
	NativeName                *string `json:"nativename"`
	NationalityCountryIsoCode *string `json:"nationality_country_iso_code"`
	Code                      *string `json:"code"`
	DateOfBirth               *string `json:"date_of_birth"`
	CountryOfBirthISOCode     *string `json:"country_of_birth_iso_code"`
	Gender                    *string `json:"gender"`
	Address                   *string `json:"address"`
	PostalCode                *string `json:"postal_code"`
	City                      *string `json:"city"`
	CountryISOCode            *string `json:"country_iso_code"`
	MSISDN                    *string `json:"msisdn"`
	Email                     *string `json:"email"`
	IDType                    *string `json:"id_type"`
	IDCountryISOCode          *string `json:"id_country_iso_code"`
	IDNumber                  *string `json:"id_number"`
	IDDeliveryDate            *string `json:"id_delivery_date"`
	IDExpirationDate          *string `json:"id_expiration_date"`
	Occupation                *string `json:"occupation"`
	Bank                      *string `json:"bank"`
	BankAccount               *string `json:"bank_account"`
	Card                      *string `json:"card"`
	ProvinceState             *string `json:"province_state"`
	BeneficiaryRelationship   *string `json:"beneficiary_relationship"`
	SourceOfFunds             *string `json:"source_of_funds"`
}

// SendingBusinessInformation represents sending business information for a given transaction of type B2C or B2B.
type SendingBusinessInformation struct {
	RegisteredName                 *string `json:"registered_name"`
	TradingName                    *string `json:"trading_name"`
	Address                        *string `json:"address"`
	PostalCode                     *string `json:"postal_code"`
	City                           *string `json:"city"`
	ProvinceState                  *string `json:"province_state"`
	CountryISOCode                 *string `json:"country_iso_code"`
	MSISDN                         *string `json:"msisdn"`
	Email                          *string `json:"email"`
	RegistrationNumber             *string `json:"registration_number"`
	Code                           *string `json:"code"`
	TaxID                          *string `json:"tax_id"`
	DateOfIncorporation            *string `json:"date_of_incorporation"`
	RepresentativeLastName         *string `json:"representative_lastname"`
	RepresentativeLastName2        *string `json:"representative_lastname2"`
	RepresentativeFirstName        *string `json:"representative_firstname"`
	RepresentativeMiddleName       *string `json:"representative_middlename"`
	RepresentativeNativeName       *string `json:"representative_nativename"`
	RepresentativeIDType           *string `json:"representative_id_type"`
	RepresentativeIDCountryISOCode *string `json:"representative_id_country_iso_code"`
	RepresentativeIDNumber         *string `json:"representative_id_number"`
	RepresentativeIDDeliveryDate   *string `json:"representative_id_delivery_date"`
	RepresentativeIDExpirationDate *string `json:"representative_id_expiration_date"`
}
