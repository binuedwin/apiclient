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

type TransactionTypes struct {
	C2C BaseTransactionInfomation `json:"C2C,omitempty"`
	B2C BaseTransactionInfomation `json:"B2C,omitempty"`
}

// TransactionTypes requires custom unmarshal due to it's dynamic json values
// func ToTransactionType(data map[string]interface{}) (TransactionType, error) {
// 	transMarshal := func(data interface{}, v interface{}) error {
// 		dataBytes, err := json.Marshal(data)
// 		if err != nil {
// 			return err
// 		}

// 		return json.Unmarshal(dataBytes, v)
// 	}

// 	if data, ok := data["C2C"]; ok {
// 		var c2c C2C
// 		if err := transMarshal(data, &c2c); err != nil {
// 			return nil, err
// 		}

// 		return TransactionType(&c2c), nil
// 	}

// 	if _, ok := data["B2C"]; ok {
// 		var b2c B2C
// 		if err := transMarshal(data, &b2c); err != nil {
// 			return nil, err
// 		}

// 		return TransactionType(&b2c), nil
// 	}

// 	return nil, errors.New("could not parse the transaction type")
// }

type C2C struct {
	BaseTransactionInfomation
}

func (ts *C2C) transactionName() string { return "C2C" }

type B2C struct {
	BaseTransactionInfomation
}

func (ts *B2C) transactionName() string { return "B2C" }

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
