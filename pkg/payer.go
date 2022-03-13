package pkg

type Payer struct {
	ID                   int                    `json:"id"`
	Name                 string                 `json:"name"`
	Precision            int                    `json:"precision,omitempty"`
	Increment            float32                `json:"increment,omitempty"`
	Currency             string                 `json:"currency"`
	CountryISOCode       string                 `json:"country_iso_code"`
	MinTransactionAmount int                    `json:"minimum_transaction_amount,omitempty"`
	MaxTransactionAmount int                    `json:"maximum_transaction_amount,omitempty"`
	Service              Service                `json:"service"`
	TransactionTypes     map[string]interface{} `json:"transaction_types,omitempty"`
}

type BaseTransactionInfomation struct {
	MinTransactionAmount            string                  `json:"minimum_transaction_amount"`
	MaxTransactionAmount            string                  `json:"maximum_transaction_amount"`
	CreditPartyIndentifiersAccepted [][]string              `json:"credit_party_identifiers_accepted"`
	RequiredSendingIdentityFields   [][]string              `json:"required_sending_entity_fields"`
	RequiredReceivingIdentityFields [][]string              `json:"required_receiving_entity_fields"`
	RequiredDocuments               [][]string              `json:"required_documents"`
	CreditPartyInformation          CreditPartyInformation  `json:"credit_party_information"`
	CreditPartyVerification         CreditPartyVerification `json:"credit_party_verification"`
}

type CreditPartyInformation struct {
	CreditPartyIdentifiersAccepted [][]string `json:"credit_party_identifiers_accepted"`
}

type CreditPartyVerification struct {
	CreditPartyIdentifiersAccepted  [][]string `json:"credit_party_identifiers_accepted"`
	RequiredReceivingIdentityFields [][]string `json:"required_receiving_entity_fields"`
}
