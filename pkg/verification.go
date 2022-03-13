package pkg

/*
AVAILABLE 							Credit party account is available and can receive a transfer
UNREGISTERED 						Credit party account is not registered but can still receive a transfer
UNAVAILABLE 						Credit party account is not available and will not receive a transfer
UNAVAILABLE-BENEFICIARY-MISMATCH 	Credit party account does not match the beneficiary details
UNAVAILABLE-INVALID-ACCOUNT 		Credit party account number is invalid
UNAVAILABLE-BARRED-ACCOUNT 			Credit party account number is barred
*/

type VerificationStatus struct {
	ID            int    `json:"id"`
	AccountStatus string `json:"account_status"`
}

type VerificationStatusRequest struct {
	CreditPartyIdentifier CreditPartyIdentifier `json:"credit_party_identifier"`
}
