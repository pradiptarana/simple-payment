package model

type AccountValidationRequest struct {
	Bank    string `form:"bank"`
	Account string `form:"account"`
}

type PayoutData struct {
	BeneficiaryName    string `json:"beneficiary_name"`
	BeneficiaryAccount string `json:"beneficiary_account"`
	BeneficiaryBank    string `json:"beneficiary_bank"`
	Amount             string `json:"amount"`
	Notes              string `json:"notes"`
}

type PayoutRequest struct {
	Payouts []PayoutData `json:"payouts"`
}

type PayoutNotificationRequest struct {
	ReferenceNo string `json:"reference_no"`
	Amount      string `json:"amount"`
	Status      string `json:"status"`
	UpdatedAt   string `json:"updated_at"`
}
