package kurs

type Currency struct{
	Nilai string `json:"nilai"`
	KursJual string `json:"kursjual"`
	KursBeli string `json:"kursbeli"`
}

type Data struct{
	LastUpdated string
	Currency map[string]Currency
}
