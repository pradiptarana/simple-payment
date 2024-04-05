package payout

import (
	"sync"
	"time"

	"github.com/pradiptarana/simple-payment/model"
	"github.com/pradiptarana/simple-payment/repository"
	"github.com/sirupsen/logrus"
)

type PayoutUC struct {
	repository.PayoutRepository
	*logrus.Logger
}

func NewPayoutUC(repo repository.PayoutRepository, log *logrus.Logger) *PayoutUC {
	return &PayoutUC{repo, log}
}

func (uc *PayoutUC) Payout(req *model.PayoutRequest) ([]*model.BankPayoutResponse, error) {
	var wg sync.WaitGroup
	payoutResponsesCh := make(chan *model.BankPayoutResponse, len(req.Payouts))
	errorsChan := make(chan error, len(req.Payouts))
	for _, payout := range req.Payouts {
		wg.Add(1)
		go func(payout model.PayoutData) {
			defer wg.Done()
			payoutData, err := uc.PayoutRepository.Payout(req)
			if err != nil {
				errorsChan <- err
			}
			err = uc.PayoutRepository.CreateTransaction(&model.Transaction{
				Amount:      payout.Amount,
				Status:      "pending",
				ReferenceNo: payoutData.ReferenceNo,
				AccountNo:   payout.BeneficiaryAccount,
				AccountName: payout.BeneficiaryName,
				Bank:        payout.BeneficiaryBank,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			})
			if err != nil {
				errorsChan <- err
			}
			payoutResponsesCh <- payoutData
		}(payout)
	}

	wg.Wait()
	close(payoutResponsesCh)
	close(errorsChan)

	payoutResponses := make([]*model.BankPayoutResponse, 0, len(req.Payouts))
	for chValue := range payoutResponsesCh {
		payoutResponses = append(payoutResponses, chValue)
	}

	for err := range errorsChan {
		uc.Logger.Error(err)
	}

	return payoutResponses, nil
}

func (uc *PayoutUC) PayoutNotification(notif *model.PayoutNotificationRequest) error {
	err := uc.PayoutRepository.UpdateTransaction(&model.Transaction{
		Status:      notif.Status,
		ReferenceNo: notif.ReferenceNo,
		Amount:      notif.Amount,
		UpdatedAt:   time.Now(),
	})
	if err != nil {
		uc.Logger.Error(err)
		return err
	}
	return nil
}

func (uc *PayoutUC) GetAccountValidation(accountReq *model.AccountValidationRequest) (*model.BankAccount, error) {
	accountData, err := uc.PayoutRepository.GetAccountValidation(accountReq)
	if err != nil {
		uc.Logger.Error(err)
		return nil, err
	}
	return accountData, nil
}
