package usecase

import "github.com/pradiptarana/simple-payment/model"

type PayoutUsecase interface {
	GetAccountValidation(*model.AccountValidationRequest) (*model.BankAccount, error)
	Payout(*model.PayoutRequest) ([]*model.BankPayoutResponse, error)
	PayoutNotification(*model.PayoutNotificationRequest) error
}
