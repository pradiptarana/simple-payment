package repository

import "github.com/pradiptarana/simple-payment/model"

//go:generate mockgen -destination=../mocks/mock_payout.go -package=mocks github.com/pradiptarana/simple-payment/repository PayoutRepository
type PayoutRepository interface {
	GetAccountValidation(*model.AccountValidationRequest) (*model.BankAccount, error)
	Payout(*model.PayoutRequest) (*model.BankPayoutResponse, error)
	CreateTransaction(*model.Transaction) error
	UpdateTransaction(*model.Transaction) error
}
