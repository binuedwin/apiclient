package main

import (
	"log"

	"thunes-client/api"
)

const (
	// BASE_URL provides the base url to which all requests are sent to the Thunes API.
	// The url also specifies the preproduction environment.
	BASE_URL   = "https://api-mt.pre.thunes.com/"
	API_KEY    = "ad4bd59f-5b33-5757-a042-ac257a2e2ebb"
	API_SECRET = "d6066c5e-ac70-5fd7-925e-8ce7751b3b69"
)

func main() {
	// construct the thunes client and make a ping request
	tc := api.NewThunesClient(BASE_URL, API_KEY, API_SECRET)

	// ping the API to make sure credentials work
	status, err := tc.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// make sure the status provided is "up"
	if status.Status != "up" {
		log.Fatalf("status returned is %s", status.Status)
	}

	// // get the list of services available
	// services, err := listServices()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("All services are: %v\n", services)

	// // get the list of all available payers from our caller account
	// // for the service id for "MobileWallet"
	// //
	// // "MobileWallet" being the only service available for our pre-production account
	// payers, err := listPayers()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("payers: %+v\n", payers)

	// var idPaidTo int
	// for _, payer := range payers {
	// 	// create a MobileWallet transaction to the first payer account found
	// 	if payer.Service.Name == "MobileWallet" {
	// 		// retrive payer's rate
	// 		rates, err := retrivePayersRate(payer.ID)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		fmt.Printf("Rates: %+v\n", rates)

	// 		// check that balance is > 150 then pay 100
	// 		balances, err := checkBalance()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}

	// 		// create a quotation
	// 		fmt.Printf("%+v\n", balances)

	// 		// quotation, err := createQuotation(100,payer.Currency, )
	// 		// if err != nil {
	// 		// 	log.Fatal(err)
	// 		// }

	// 		// introspect the error
	// 		//fmt.Printf("%+v\n", quotation)

	// 		break
	// 	}
	// }

	// fmt.Printf("Paid 100 to: %d", idPaidTo)
}
