package pkg

type PayerRates struct {
	DestinationCurrency string                        `json:"destination_currency"`
	Rates               map[string]map[string][]Rates `json:"rates"`
}

type Rates struct {
	SourceAmountMin int64   `json:"source_amount_min"`
	SourceAmountMax int64   `json:"source_amount_max"`
	WholesaleFXRate float64 `json:"wholesale_fx_rate"`
}
