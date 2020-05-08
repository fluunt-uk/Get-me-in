package models

type CardDetails struct {
	Number		string `json:"card_number"`
	ExpMonth	string `json:"expiry_month"`
	ExpYear		string `json:"expiry_year"`
	CVC			string `json:"cvc"`
}

