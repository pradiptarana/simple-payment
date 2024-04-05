package model

import "time"

type BankAccount struct {
	Bank          string `json:"bank_name"`
	AccountName   string `json:"account_name"`
	AccountNumber string `json:"account_no"`
}

type Transaction struct {
	Id          int       `json:"id"`
	Amount      string    `json:"amount"`
	Status      string    `json:"status"`
	ReferenceNo string    `json:"reference_no"`
	AccountNo   string    `json:"account_no"`
	AccountName string    `json:"account_name"`
	Bank        string    `json:"bank"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type URL struct {
	Id          int    `db:"id"`
	OriginalURL string `db:"original_url"`
	ShortURL    string `db:"short_url"`
	CreatedAt   int64  `db:"created_at"`
}

type BankPayoutResponse struct {
	ReferenceNo string `json:"reference_no"`
	Status      string `json:"status"`
}
