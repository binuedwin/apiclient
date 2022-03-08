package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"thunes-client/pkg"
)

// ping is used to check the validity of the credentials we need to interact with the
// Thunes API
func (tc *ThunesClient) Ping() (*pkg.Status, error) {
	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, "ping", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var status pkg.Status
	if err = json.Unmarshal(dataOut, &status); err != nil {
		return nil, err
	}

	return &status, nil
}

// listServices is used to check the available services that we can access with our
// caller account
func (tc *ThunesClient) ListServices() ([]pkg.Service, error) {
	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, "v2/money-transfer/services", nil, http.StatusOK)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var services []pkg.Service
	if err = json.Unmarshal(dataOut, &services); err != nil {
		return nil, err
	}

	return services, nil
}

// listPayers provides a list of all available payers for our caller account
// for the service id provided
func (tc *ThunesClient) ListPayers(per_page *int) ([]pkg.Payer, error) {
	if per_page == nil {
		*per_page = 100
	}

	// construct the request
	dataOut, err := tc.NewRequest(
		http.MethodGet,
		fmt.Sprintf("v2/money-transfer/payers?per_page=%d", *per_page),
		nil,
		http.StatusOK,
	)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var payers []pkg.Payer
	if err = json.Unmarshal(dataOut, &payers); err != nil {
		return nil, err
	}

	return payers, nil
}

func (tc *ThunesClient) ListCountriesAvailable() ([]pkg.Country, error) {
	dataOut, err := tc.NewRequest(
		http.MethodGet,
		"v2/money-transfer/countries",
		nil,
		http.StatusOK,
	)

	if err != nil {
		return nil, err
	}

	// handle the success case response
	var countries []pkg.Country
	if err = json.Unmarshal(dataOut, &countries); err != nil {
		return nil, err
	}

	return countries, nil
}

func (tc *ThunesClient) BICCodeLookup(swiftBICCode string) ([]pkg.Lookup, error) {
	dataOut, err := tc.NewRequest(
		http.MethodGet,
		fmt.Sprintf("v2/money-transfer/lookups/BIC/%s", swiftBICCode),
		nil,
		http.StatusOK,
	)

	if err != nil {
		return nil, err
	}

	// handle the success case response
	var lookups []pkg.Lookup
	if err = json.Unmarshal(dataOut, &lookups); err != nil {
		return nil, err
	}

	return lookups, nil
}

// Returns
// Beneficiary object if C2C or B2C transaction type.
// Receiving business information object if C2B or B2B transaction type.
func (tc *ThunesClient) CreditPartyInformation(id, transactionType string) (interface{}, error) {
	dataOut, err := tc.NewRequest(
		http.MethodGet,
		fmt.Sprintf("v2/money-transfer/payers/%s/%s/credit-party-information", id, transactionType),
		nil,
		http.StatusOK,
	)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	if transactionType == "C2C" || transactionType == "B2C" {
		var beneficiary pkg.Beneficiary
		if err = json.Unmarshal(dataOut, &beneficiary); err != nil {
			return nil, err
		}

		return beneficiary, nil
	}

	//else if transactionType == "C2B" || transactionType == "B2B"
	var receivingBusinessInfo pkg.ReceivingBusinessInformation
	if err = json.Unmarshal(dataOut, &receivingBusinessInfo); err != nil {
		return nil, err
	}

	return receivingBusinessInfo, nil
}

func (tc *ThunesClient) CreditPartyVerification(id int, mssidn, transactionType string) (*pkg.VerificationStatus, error) {
	// construct the request body
	verificationStatusReq := pkg.VerificationStatusRequest{
		CreditPartyIdentifier: pkg.CreditPartyIdentifier{
			MSISDN: mssidn,
		},
	}
	data, err := json.Marshal(verificationStatusReq)
	if err != nil {
		return nil, err
	}

	// construct the request
	dataOut, err := tc.NewRequest(
		http.MethodPost,
		fmt.Sprintf("v2/money-transfer/payers/%d/%s/credit-party-verification", id, transactionType),
		bytes.NewBuffer(data),
		http.StatusOK,
	)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var verStatus pkg.VerificationStatus
	if err = json.Unmarshal(dataOut, &verStatus); err != nil {
		return nil, err
	}

	return &verStatus, nil
}

func (tc *ThunesClient) RetrivePayersRate(id int) (*pkg.PayerRates, error) {
	// construct the request
	dataOut, err := tc.NewRequest(
		http.MethodGet,
		fmt.Sprintf("v2/money-transfer/payers/%d/rates", id),
		nil,
		http.StatusOK,
	)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var rates pkg.PayerRates
	if err = json.Unmarshal(dataOut, &rates); err != nil {
		return nil, err
	}

	return &rates, nil
}

func (tc *ThunesClient) CreateQuotation(amount int, destinationCurrency, currency, payeyID, countryISOCode string) (*pkg.Quotation, error) {
	// quotation request body
	amt := int64(amount)
	quotationReq := pkg.CreateQuotationRequest{
		ExternalID:      "some-long-id",
		PayerID:         payeyID,
		Mode:            "SOURCE_AMOUNT",
		TransactionType: "C2C",
		Source: pkg.Source{
			Amount:         &amt,
			Currency:       currency,
			CountryISOCode: countryISOCode,
		},
		Destination: pkg.CurrencyAmount{
			Currency: destinationCurrency,
			Amount:   nil,
		},
	}

	data, err := json.Marshal(quotationReq)
	if err != nil {
		return nil, err
	}

	// construct the request
	dataOut, err := tc.NewRequest(
		http.MethodPost,
		"v2/money-transfer/quotations",
		bytes.NewBuffer(data),
		http.StatusCreated,
	)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var quotation pkg.Quotation
	if err = json.Unmarshal(dataOut, &quotation); err != nil {
		return nil, err
	}

	return &quotation, nil
}

func (tc *ThunesClient) CheckBalance(per_page *int) ([]pkg.Balance, error) {
	if per_page == nil {
		*per_page = 100
	}

	// construct the request
	dataOut, err := tc.NewRequest(
		http.MethodGet,
		fmt.Sprintf("v2/money-transfer/balances?per_page=%d", *per_page),
		nil,
		http.StatusOK,
	)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var balances []pkg.Balance
	if err = json.Unmarshal(dataOut, &balances); err != nil {
		return nil, err
	}

	return balances, nil
}


