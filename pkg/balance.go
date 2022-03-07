package pkg

type Balance struct {
	ID             int    `json:"id"`
	Currency       string `json:"currency"`
	Balance        int64  `json:"balance"`
	Pending        int64  `json:"pending"`
	Available      int64  `json:"available"`
	CreditFacility int64  `json:"credit_facility"`
}
