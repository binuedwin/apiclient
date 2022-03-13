package pkg

// Transaction represents transaction information for a transfer request.
type Transaction struct {
	ID                        *int                          `json:"id"`
	Status                    *string                       `json:"status"`
	StatusMessage             *string                       `json:"status_message"`
	StatusClass               *string                       `json:"status_class"`
	StatusClassMessage        *string                       `json:"status_class_message"`
	ExternalID                *string                       `json:"external_id"`
	ExternalCode              *string                       `json:"external_code"`
	TransactionType           *string                       `json:"transaction_type"`
	PayerTransactionReference *string                       `json:"payer_transaction_reference"`
	PayerTransactionCode      *string                       `json:"payer_transaction_code"`
	CreationDate              *string                       `json:"creation_date"`
	ExpirationDate            *string                       `json:"expiration_date"`
	CreditPartyIdentifier     *CreditPartyIdentifier        `json:"credit_party_identifier"`
	Source                    *Source                       `json:"source"`
	Destination               *CurrencyAmount               `json:"destination"`
	Payer                     *Payer                        `json:"payer"`
	Sender                    *Sender                       `json:"sender"`
	Beneficiary               *Beneficiary                  `json:"beneficiary"`
	SendingBusiness           *SendingBusinessInformation   `json:"sending_business"`
	ReceivingBusiness         *ReceivingBusinessInformation `json:"receiving_business"`
	CallbackURL               *string                       `json:"callback_url"`
	SendAmount                *CurrencyAmount               `json:"send_amount"`
	WholeSaleFXRate           *float64                      `json:"wholesale_fx_rate"`
	RetailRate                *float64                      `json:"retail_rate"`
	RetailFee                 *float64                      `json:"retail_fee"`
	RetailFeeCurrency         *string                       `json:"retail_fee_currency"`
	Fee                       *CurrencyAmount               `json:"fee"`
	PurposeOfRemittance       *string                       `json:"purpose_of_remittance"`
	DocumentReferenceNumber   *string                       `json:"document_reference_number"`
	AdditionalInformation1    *string                       `json:"additional_information_1"`
	AdditionalInformation2    *string                       `json:"additional_information_2"`
	AdditionalInformation3    *string                       `json:"additional_information_3"`
}

// CreateTransactionRequest represents the request body for creating a transaction.
//
// Requiredness will depend on the chosen transaction type, quotation of type :
//		C2C will require a sender and a beneficiary
//		C2B will require a sender and a receiving business
//		B2C will require a sending business and a beneficiary
//		B2B will require a sending business and a receiving business
type CreateTransactionRequest struct {
	CreditPartyIdentifier   *CreditPartyIdentifier        `json:"credit_party_identifier"`   // mandatory
	RetailFee               *float64                      `json:"retail_fee"`                // optional
	RetailRate              *float64                      `json:"retail_rate"`               // optional
	RetailFeeCurrency       *string                       `json:"retail_fee_currency"`       // optional
	Sender                  *Sender                       `json:"sender"`                    // depends on transaction type
	Beneficiary             *Beneficiary                  `json:"beneficiary"`               // depends on transaction type
	SendingBusiness         *SendingBusinessInformation   `json:"sending_business"`          // depends on transaction type
	ReceivingBusiness       *ReceivingBusinessInformation `json:"receiving_business"`        // depends on transaction type
	ExternalID              *string                       `json:"external_id"`               // mandatory
	ExternalCode            *string                       `json:"external_code"`             // optional
	CallbackURL             *string                       `json:"callback_url"`              // optional
	PurposeOfRemittance     *string                       `json:"purpose_of_remittance"`     // B2B: Yes, Other: No
	DocumentReferenceNumber *string                       `json:"document_reference_number"` // B2B: Yes, Other: No
	AdditionalInformation1  *string                       `json:"additional_information_1"`  // optional
	AdditionalInformation2  *string                       `json:"additional_information_2"`  // optional
	AdditionalInformation3  *string                       `json:"additional_information_3"`  // optional
}

// TransactionAttachment represents an attachment which has been added to a transaction.
type TransactionAttachment struct {
	ID            int    `json:"id"`
	TransactionID int    `json:"transaction_id"`
	Name          string `json:"name"`
	ContentType   string `json:"content_type"`
	Type          string `json:"type"`
}

type TransactionAttachmentType string

const (
	INVOICE        = TransactionAttachmentType("invoice")
	PURCHASE_ORDER = TransactionAttachmentType("purchase_order")
	DELIVERY_SLIP  = TransactionAttachmentType("delivery_slip")
	CONTRACT       = TransactionAttachmentType("contract")
)
