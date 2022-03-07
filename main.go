package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"thunes-client/pkg"
)

const (
	// BASE_URL provides the base url to which all requests are sent to the Thunes API.
	// The url also specifies the preproduction environment.
	BASE_URL   = "https://api-mt.pre.thunes.com/"
	API_KEY    = "ad4bd59f-5b33-5757-a042-ac257a2e2ebb"
	API_SECRET = "d6066c5e-ac70-5fd7-925e-8ce7751b3b69"
)

func main() {
	// ping the API to make sure credentials work
	status, err := ping()
	if err != nil {
		log.Fatal(err)
	}

	// make sure the status provided is "up"
	if status.Status != "up" {
		log.Fatal(fmt.Sprintf("status returned is %s", status.Status))
	}

	// get the list of services available
	services, err := listServices()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("All services are: %v\n", services)

	// get the list of all available payers from our caller account
	// for the service id for "MobileWallet"
	//
	// "MobileWallet" being the only service available for our pre-production account
	payers, err := listPayers()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", payers)

	var idPaidTo int
	for _, payer := range payers {
		// create a MobileWallet transaction to the first payer account found
		if payer.Service.Name == "MobileWallet" {
			// check that balance is > 150 then pay 100
			balances, err := checkBalance()
			if err != nil {
				log.Fatal(err)
			}

			// create a quotation
			fmt.Printf("%+v\n", balances)

			quotation, err := createQuotation(100)
			if err != nil {
				log.Fatal(err)
			}

			// introspect the error
			fmt.Printf("%+v\n", quotation)

			break
		}
	}

	fmt.Printf("Paid 100 to: %d", idPaidTo)
}

func createQuotation(amount int) (*pkg.Quotation, error) {
	// quotation request body

	quotationReq := pkg.CreateQuotationRequest{
		ExternalID:      "some-long-id",
		PayerID:         "1",
		Mode:            "SOURCE_AMOUNT",
		TransactionType: "C2C",
		Source: pkg.Source{
			Amount:         int64(amount),
			Currency:       "UAE",
			CountryISOCode: "AED",
		},
		Destination: pkg.CurrencyAmount{
			Currency: "UAE",
			Amount:   nil,
		},
	}

	data, err := json.Marshal(quotationReq)
	if err != nil {
		return nil, err
	}

	// construct the request
	req, err := http.NewRequest(
		http.MethodPost,
		BASE_URL+"v2/money-transfer/quotations",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}

	// set the auth
	req.SetBasicAuth(API_KEY, API_SECRET)

	// make the http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// handle the error response case
	if resp.StatusCode != http.StatusCreated {
		return nil, respError(resp.Body)
	}

	// handle the success case response
	var quotation pkg.Quotation
	if err = json.NewDecoder(resp.Body).Decode(&quotation); err != nil {
		return nil, err
	}

	return &quotation, nil
}

func checkBalance() ([]pkg.Balance, error) {
	// construct the request
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%sv2/money-transfer/balances?per_page=100", BASE_URL),
		nil,
	)
	if err != nil {
		return nil, err
	}

	// set the auth
	req.SetBasicAuth(API_KEY, API_SECRET)

	// make the http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// handle the error response case
	if resp.StatusCode != http.StatusOK {
		return nil, respError(resp.Body)
	}

	// handle the success case response
	var balances []pkg.Balance
	if err = json.NewDecoder(resp.Body).Decode(&balances); err != nil {
		return nil, err
	}

	return balances, nil
}

// listPayers provides a list of all available payers for our caller account
// for the service id provided
func listPayers() ([]pkg.Payer, error) {
	// construct the request
	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%sv2/money-transfer/payers?per_page=100", BASE_URL),
		nil,
	)
	if err != nil {
		return nil, err
	}

	// set the auth
	req.SetBasicAuth(API_KEY, API_SECRET)

	// make the http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// handle the error response case
	if resp.StatusCode != http.StatusOK {
		return nil, respError(resp.Body)
	}

	// handle the success case response
	var payers []pkg.Payer
	if err = json.NewDecoder(resp.Body).Decode(&payers); err != nil {
		return nil, err
	}

	return payers, nil
}

// listServices is used to check the available services that we can access with our
// caller account
func listServices() ([]pkg.Service, error) {
	// construct the request
	req, err := http.NewRequest(http.MethodGet, BASE_URL+"v2/money-transfer/services", nil)
	if err != nil {
		return nil, err
	}

	// set the auth
	req.SetBasicAuth(API_KEY, API_SECRET)

	// make the http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// handle the error response case
	if resp.StatusCode != http.StatusOK {
		return nil, respError(resp.Body)
	}

	// handle the success case response
	var services []pkg.Service
	if err = json.NewDecoder(resp.Body).Decode(&services); err != nil {
		return nil, err
	}

	return services, nil
}

// ping is used to check the validity of the credentials we need to interact with the
// Thunes API
func ping() (*pkg.Status, error) {
	// construct the request
	req, err := http.NewRequest(http.MethodGet, BASE_URL+"ping", nil)
	if err != nil {
		return nil, err
	}

	// set the auth
	req.SetBasicAuth(API_KEY, API_SECRET)

	// make the http request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// handle the error response case
	if resp.StatusCode != http.StatusOK {
		return nil, respError(req.Body)
	}

	// handle the success case response
	var status pkg.Status
	if err = json.NewDecoder(resp.Body).Decode(&status); err != nil {
		return nil, err
	}

	return &status, nil
}

func respError(body io.Reader) error {
	var errors pkg.Errors
	if err := json.NewDecoder(body).Decode(&errors); err != nil {
		return err
	}

	return errors.ToError()
}
