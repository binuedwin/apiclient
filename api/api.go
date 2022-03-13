package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"thunes-client/pkg"
)

/*
Transaction attachment allowed file types:
.txt 	Plain Text 	text/plain
.pdf 	Adobe Portable Document Format 	application/pdf
.doc 	Microsoft Word Document 	application/msword
.docx 	Microsoft Word (OpenXML) 	application/vnd.openxmlformats-officedocument.wordprocessingml.document
.jpg, .jpeg 	JPEG images 	image/jpeg
.png 	Portable Network Graphics 	image/png
.bmp 	Windows OS/2 Bitmap Graphics 	image/bmp
.rtf 	Rich Text Format (RTF) 	application/rtf
.xls 	Microsoft Excel 	application/vnd.ms-excel
.xlsx 	Microsoft Excel (OpenXML) 	application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
*/
var TransactionAttachmentAllowed = []string{
	".txt", ".pdf", ".doc", ".docx", ".jpg", ".jpeg", ".png", ".bmp", ".rtf", ".xls", ".xlsx",
}

// Ping is used to check the validity of the credentials we need to interact with the Thunes API.
func (tc *ThunesClient) Ping() (*pkg.Status, error) {
	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, "ping", nil, http.StatusOK, nil)
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

// ListServices is used to check the available services that we can access with our caller account.
// The arguements to this method are optional, nil refrences can be passed.
func (tc *ThunesClient) ListServices(page, perPage *int, countryISOCode *string) ([]pkg.Service, error) {
	// construct query params if they are present
	queryParams := make(map[string]string)
	if page != nil {
		queryParams["page"] = fmt.Sprint(*page)
	}
	if perPage != nil {
		queryParams["per_page"] = fmt.Sprint(*perPage)
	}
	if countryISOCode != nil {
		queryParams["country_iso_code"] = *countryISOCode
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, "v2/money-transfer/services", nil, http.StatusOK, queryParams)
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

// ListPayers provides a list of all available payers for our caller account for the service id provided.
// The arguements to this method are optional, nil references can be passed.
func (tc *ThunesClient) ListPayers(page, perPage, serviceID *int, countryISOCode, currency *string) ([]pkg.Payer, error) {
	// construct query params if they are present
	queryParams := make(map[string]string)
	if page != nil {
		queryParams["page"] = fmt.Sprint(*page)
	}
	if perPage != nil {
		queryParams["per_page"] = fmt.Sprint(*perPage)
	}
	if serviceID != nil {
		queryParams["service_id"] = fmt.Sprint(*serviceID)
	}
	if countryISOCode != nil {
		queryParams["country_iso_code"] = *countryISOCode
	}
	if currency != nil {
		queryParams["currency"] = *currency
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, "v2/money-transfer/payers", nil, http.StatusOK, queryParams)
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

// GetPayerDetails retrives information for a given payer.
func (tc *ThunesClient) GetPayerDetails(id int) (*pkg.Payer, error) {
	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, fmt.Sprintf("v2/money-transfer/payers/%d", id), nil, http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var payer pkg.Payer
	if err = json.Unmarshal(dataOut, &payer); err != nil {
		return nil, err
	}

	return &payer, nil
}

// GetPayerRates retrives rates for a given payer.
func (tc *ThunesClient) GetPayerRates(id int) (*pkg.PayerRates, error) {
	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, fmt.Sprintf("v2/money-transfer/payers/%d/rates", id), nil, http.StatusOK, nil)
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

// ListCountriesAvailable is used to retrieve the list of countries for all money transfer services available for the caller.
// The arguements to this method are optional, nil references can be passed.
func (tc *ThunesClient) ListCountriesAvailable(page, perPage *int) ([]pkg.Country, error) {
	// construct query params if they are present
	queryParams := make(map[string]string)
	if page != nil {
		queryParams["page"] = fmt.Sprint(*page)
	}
	if perPage != nil {
		queryParams["per_page"] = fmt.Sprint(*perPage)
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, "v2/money-transfer/countries", nil, http.StatusOK, queryParams)
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

// BICCodeLookup retrives a list of payers' identifier for a given SWIFT BIC code.
// In the arguements the swiftBICCode is mandatory while page and perPage are optional.
func (tc *ThunesClient) BICCodeLookup(swiftBICCode string, page, perPage *int) ([]pkg.Lookup, error) {
	// construct query params if they are present
	queryParams := make(map[string]string)
	if page != nil {
		queryParams["page"] = fmt.Sprint(*page)
	}
	if perPage != nil {
		queryParams["per_page"] = fmt.Sprint(*perPage)
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, fmt.Sprintf("v2/money-transfer/lookups/BIC/%s", swiftBICCode), nil, http.StatusOK, queryParams)
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

// GetBalances retrives information for all account balances per currency.
// The arguements to this method are optional, nil references can be passed.
// Formulae: available = balance - pending + credit_facility
func (tc *ThunesClient) GetBalances(page, perPage *int) ([]pkg.Balance, error) {
	// construct query params if they are present
	queryParams := make(map[string]string)
	if page != nil {
		queryParams["page"] = fmt.Sprint(*page)
	}
	if perPage != nil {
		queryParams["per_page"] = fmt.Sprint(*perPage)
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, "v2/money-transfer/balances", nil, http.StatusOK, queryParams)
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

// GetCreditPartyInformation retrieves Beneficiary or ReceivingBusinessInformation information based on account details for a given payer and transaction type.
// All arguements to the method are mandatory.
// Beneficiary is returned if the transactionType supplies is "C2C" or "B2C".
// ReceivingBusinessInformation is returned if the transaction type supplied is "C2B" or "B2B".
func (tc *ThunesClient) GetCreditPartyInformation(id, transactionType, msisdn string, out interface{}) error {
	// construct request body
	creditPartyInfoReq := pkg.CreditPartyIdentifierRequestWrapper{
		CreditPartyIdentifier: pkg.CreditPartyIdentifier{
			MSISDN: msisdn,
		},
	}
	data, err := json.Marshal(creditPartyInfoReq)
	if err != nil {
		return err
	}

	// construct the request
	dataOut, err := tc.NewRequest(
		http.MethodGet,
		fmt.Sprintf("v2/money-transfer/payers/%s/%s/credit-party-information", id, transactionType),
		bytes.NewBuffer(data),
		http.StatusOK,
		nil,
	)
	if err != nil {
		return err
	}

	// handle the success case response
	if err = json.Unmarshal(dataOut, &out); err != nil {
		return err
	}

	return nil
}

// CreditPartyVerification validates the status of an account for a given payer and transaction type.
// All arguement to be supplied are mandatory.
func (tc *ThunesClient) CreditPartyVerification(id int, mssidn, transactionType string) (*pkg.VerificationStatus, error) {
	// construct the request body
	verificationStatusReq := pkg.CreditPartyIdentifierRequestWrapper{
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
		nil,
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

// CreateQuotationForSource creates a new quotation for a source value.
// All arguements to be supplied are mandatory.
func (tc *ThunesClient) CreateQuotationForSource(sourceAmt int, destinationCurrency, sourceCurrency, payerID, sourceCountryISOCode, transactionType, externalID string) (*pkg.Quotation, error) {
	// quotation request body
	amt := int64(sourceAmt)
	quotationReq := pkg.CreateQuotationRequest{
		ExternalID:      externalID,
		PayerID:         payerID,
		Mode:            "SOURCE_AMOUNT",
		TransactionType: transactionType,
		Source: pkg.Source{
			Amount:         &amt,
			Currency:       sourceCurrency,
			CountryISOCode: sourceCountryISOCode,
		},
		Destination: pkg.CurrencyAmount{
			Currency: destinationCurrency,
			Amount:   nil,
		},
	}

	return tc.createQuotation(&quotationReq)
}

// CreateQuotationForDestination creates a new quotation for a destination value.
// All arguements to be supplied are mandatory.
func (tc *ThunesClient) CreateQuotationForDestination(destinationAmt int, destinationCurrency, sourceCurrency, payerID, sourceCountryISOCode, transactionType, externalID string) (*pkg.Quotation, error) {
	// quotation request body
	amt := int64(destinationAmt)
	quotationReq := pkg.CreateQuotationRequest{
		ExternalID:      externalID,
		PayerID:         payerID,
		Mode:            "DESTINATION_AMOUNT",
		TransactionType: transactionType,
		Source: pkg.Source{
			Amount:         nil,
			Currency:       sourceCurrency,
			CountryISOCode: sourceCountryISOCode,
		},
		Destination: pkg.CurrencyAmount{
			Currency: destinationCurrency,
			Amount:   &amt,
		},
	}

	return tc.createQuotation(&quotationReq)
}

func (tc *ThunesClient) createQuotation(quotationReq *pkg.CreateQuotationRequest) (*pkg.Quotation, error) {
	// parse to json
	data, err := json.Marshal(quotationReq)
	if err != nil {
		return nil, err
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodPost, "v2/money-transfer/quotations", bytes.NewBuffer(data), http.StatusCreated, nil)
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

// GetQuotationByID retrieves information for a given quotation given the quotation id.
func (tc *ThunesClient) GetQuotationByID(id int) (*pkg.Quotation, error) {
	// construct the request
	dataOut, err := tc.NewRequest(http.MethodPost, fmt.Sprintf("v2/money-transfer/quotations/%d", id), nil, http.StatusOK, nil)
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

// GetQuotationByExternalID retrieves information for a given quotation given the quotation externalID.
func (tc *ThunesClient) GetQuotationByExternalID(externalID string) (*pkg.Quotation, error) {
	// construct the request
	dataOut, err := tc.NewRequest(http.MethodPost, fmt.Sprintf("v2/money-transfer/quotations/ext-%s", externalID), nil, http.StatusOK, nil)
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

// CreateTransaction create a new transaction with transfer values specified from a given quotation.
// Either the ID or the externalID of the quotation must be supplied. If both are supplied, the ID will be used.
func (tc *ThunesClient) CreateTransaction(reqBody *pkg.CreateTransactionRequest, id *int, externalID *string) (*pkg.Transaction, error) {
	if id == nil && externalID == nil {
		return nil, errors.New("either the ID or the externalID of the quotation must be supplied")
	}

	var reqURL string
	if id != nil {
		reqURL = fmt.Sprintf("v2/money-transfer/quotations/%d/transactions", *id)
	} else {
		reqURL = fmt.Sprintf("v2/money-transfer/quotations/ext-%s/transactions", *externalID)
	}

	// parse to json
	data, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodPost, reqURL, bytes.NewBuffer(data), http.StatusCreated, nil)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var transaction pkg.Transaction
	if err = json.Unmarshal(dataOut, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

// AddAttachmentToTransaction adds an attachemnt to a given transaction.
// There is a maximum of 3 files that can be sent per transaction.
// Either the ID or the externalID of the transaction must be supplied. If both are supplied, the ID will be used.
func (tc *ThunesClient) AddAttachmentToTransaction(name, transactionAttachmentType pkg.TransactionAttachmentType, file *os.File, id *int, externalID *string) (*pkg.TransactionAttachment, error) {
	// ensure the transaction ID or externalID is supplied
	if id == nil && externalID == nil {
		return nil, errors.New("either the ID or the externalID of the transaction must be supplied")
	}

	fInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	// if file extension is not allowed in the attachment type, return an error
	if !containedIn(TransactionAttachmentAllowed, filepath.Ext(file.Name())) {
		return nil, errors.New("file extension is not allowed in the attachment type")
	}

	// ensure that maximum size is 8MB
	if fInfo.Size() > 8*1024*1024 {
		return nil, errors.New("maximum file size is 8MB")
	}

	var reqURL string
	if id != nil {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/%d/attachments", *id)
	} else {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/ext-%s/attachments", *externalID)
	}

	// construct the formdata
	var reqBody bytes.Buffer
	form := multipart.NewWriter(&reqBody)
	defer form.Close()

	// add the "type" field
	{
		typeFormField, err := form.CreateFormField("type")
		if err != nil {
			return nil, err
		}

		if _, err = typeFormField.Write([]byte(string(transactionAttachmentType))); err != nil {
			return nil, err
		}
	}

	// add the file to the form with "name" field
	{
		part, err := form.CreateFormFile(file.Name(), file.Name())
		if err != nil {
			return nil, err
		}

		if _, err = io.Copy(part, file); err != nil {
			return nil, err
		}
	}

	// construct the request
	req, err := http.NewRequest(http.MethodGet, tc.baseUrl+reqURL, &reqBody)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", form.FormDataContentType())

	// send the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		return nil, respError(resp.Body)
	}

	// handle the success case response
	var transactionAttachment pkg.TransactionAttachment
	if err := json.NewDecoder(resp.Body).Decode(&transactionAttachment); err != nil {
		return nil, err
	}

	return &transactionAttachment, nil
}

// ConfirmTransaction confirms a previously-created transaction to initiate processing.
// Either the ID or the externalID of the transaction must be supplied. If both are supplied, the ID will be used.
func (tc *ThunesClient) ConfirmTransaction(id *int, externalID *string) (*pkg.Transaction, error) {
	if id == nil && externalID == nil {
		return nil, errors.New("either the ID or the externalID of the transaction must be supplied")
	}

	var reqURL string
	if id != nil {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/%d/confirm", *id)
	} else {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/ext-%s/confirm", *externalID)
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodPost, reqURL, nil, http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var transaction pkg.Transaction
	if err = json.Unmarshal(dataOut, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

// GetTransactionInformation retrives information about a given transaction.
// Either the ID or the externalID of the transaction must be supplied. If both are supplied, the ID will be used.
func (tc *ThunesClient) GetTransactionInformation(id *int, externalID *string) (*pkg.Transaction, error) {
	if id == nil && externalID == nil {
		return nil, errors.New("either the ID or the externalID of the transaction must be supplied")
	}

	var reqURL string
	if id != nil {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/%d", *id)
	} else {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/ext-%s", *externalID)
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, reqURL, nil, http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var transaction pkg.Transaction
	if err = json.Unmarshal(dataOut, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

// ListTransactionAttachments retrieves a list of all attachments for a given transaction.
func (tc *ThunesClient) ListTransactionAttachments(id *int, externalID *string) ([]pkg.TransactionAttachment, error) {
	if id == nil && externalID == nil {
		return nil, errors.New("either the ID or the externalID of the transaction must be supplied")
	}

	var reqURL string
	if id != nil {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/%d/attachments", *id)
	} else {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/ext-%s/attachments", *externalID)
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodGet, reqURL, nil, http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var transactionAttachments []pkg.TransactionAttachment
	if err = json.Unmarshal(dataOut, &transactionAttachments); err != nil {
		return nil, err
	}

	return transactionAttachments, nil
}

// CancelTransaction can be used when sending a cash pickup transaction, and a transaction is in status CONFIRMED-WAITING-FOR-PICKUP;
// a partner can call the API to cancel the previously-created transaction.
//
// If the action can’t be performed, an error 1007014 (Transaction can not be cancelled) will be returned.
// Either the ID or the externalID of the transaction must be supplied. If both are supplied, the ID will be used.
func (tc *ThunesClient) CancelTransaction(id *int, externalID *string) (*pkg.Transaction, error) {
	if id == nil && externalID == nil {
		return nil, errors.New("either the ID or the externalID of the transaction must be supplied")
	}

	var reqURL string
	if id != nil {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/%d/cancel", *id)
	} else {
		reqURL = fmt.Sprintf("v2/money-transfer/transactions/ext-%s/cancel", *externalID)
	}

	// construct the request
	dataOut, err := tc.NewRequest(http.MethodPost, reqURL, nil, http.StatusOK, nil)
	if err != nil {
		return nil, err
	}

	// handle the success case response
	var transaction pkg.Transaction
	if err = json.Unmarshal(dataOut, &transaction); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func containedIn(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
