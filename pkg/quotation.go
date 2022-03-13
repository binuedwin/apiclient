package pkg

type CreateQuotationRequest struct {
	ExternalID      string         `json:"external_id"`
	PayerID         string         `json:"payer_id"`
	Mode            string         `json:"mode"`
	TransactionType string         `json:"transaction_type"`
	Source          Source         `json:"source"`
	Destination     CurrencyAmount `json:"destination"`
}

type Quotation struct {
	ID              int            `json:"id"`
	ExternalID      string         `json:"external_id"`
	Payer           Payer          `json:"payer"`
	Mode            string         `json:"mode"`
	TransactionType string         `json:"transaction_type"`
	Source          Source         `json:"Source"`
	Destination     CurrencyAmount `json:"Destination"`
	SentAmount      CurrencyAmount `json:"sent_amount"`
	WholeSaleFXRate float64        `json:"wholesale_fx_rate"`
	Fee             CurrencyAmount `json:"fee"`
	CreationDate    string         `json:"creation_date"`
	ExpirationDate  string         `json:"expiration_date"`
}

type Source struct {
	CountryISOCode string `json:"country_iso_code"`
	Currency       string `json:"currency"`
	Amount         *int64 `json:"amount"`
}

type CurrencyAmount struct {
	Amount   *int64 `json:"amount"`
	Currency string `json:"currency"`
}
