package payout_test

import (
	"testing"

	"github.com/pradiptarana/simple-payment/mocks"
	"github.com/pradiptarana/simple-payment/model"

	"github.com/golang/mock/gomock"

	logger "github.com/pradiptarana/simple-payment/internal/logger"
	payoutUC "github.com/pradiptarana/simple-payment/usecase/payout"
)

var log = logger.Init()

func TestCreateFailedValidation(t *testing.T) {
	req := &model.PayoutRequest{}
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockPayoutRepo := mocks.NewMockPayoutRepository(mockCtrl)
	payoutUC := payoutUC.NewPayoutUC(mockPayoutRepo, log)

	_, err := payoutUC.Payout(req)
	if err == nil {
		t.Fail()
	}

	if err == nil {
		t.Fail()
	}
}

// func TestCreateFailedGetMaxId(t *testing.T) {
// 	req := &model.Payout{
// 		OriginalPayout: "https://www.google.com/abc",
// 	}
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockPayoutRepo := mocks.NewMockPayoutRepository(mockCtrl)
// 	payoutUC := payoutUC.NewPayoutUC(mockPayoutRepo, log)

// 	mockPayoutRepo.EXPECT().GetMaxID().Return(int64(0), errors.New("error get Max ID")).Times(1)

// 	_, err := payoutUC.CreateShortPayout(req)
// 	if err == nil {
// 		t.Fail()
// 	}
// }

// func TestCreateFailedCreateToDB(t *testing.T) {
// 	req := &model.Payout{
// 		OriginalPayout: "https://www.google.com/abc",
// 	}
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockPayoutRepo := mocks.NewMockPayoutRepository(mockCtrl)
// 	payoutUC := payoutUC.NewPayoutUC(mockPayoutRepo, log)

// 	mockPayoutRepo.EXPECT().GetMaxID().Return(int64(0), nil).Times(1)
// 	repoReq := req
// 	repoReq.ShortPayout = "27qMi57J"
// 	mockPayoutRepo.EXPECT().CreatePayout(repoReq).Return(errors.New("error db")).Times(1)

// 	_, err := payoutUC.CreateShortPayout(req)
// 	if err == nil {
// 		t.Fail()
// 	}
// }

// func TestCreateSuccess(t *testing.T) {
// 	req := &model.Payout{
// 		OriginalPayout: "https://www.google.com/abc",
// 	}
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockPayoutRepo := mocks.NewMockPayoutRepository(mockCtrl)
// 	payoutUC := payoutUC.NewPayoutUC(mockPayoutRepo, log)

// 	mockPayoutRepo.EXPECT().GetMaxID().Return(int64(0), nil).Times(1)
// 	repoReq := req
// 	repoReq.ShortPayout = "2XNGAK"
// 	mockPayoutRepo.EXPECT().CreatePayout(repoReq).Return(nil).Times(1)

// 	shortPayout, err := payoutUC.CreateShortPayout(req)
// 	if err != nil {
// 		t.Fail()
// 	}

// 	if shortPayout != repoReq.ShortPayout {
// 		t.Fail()
// 	}
// }

// func TestGetOriginalUrlFailed(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockPayoutRepo := mocks.NewMockPayoutRepository(mockCtrl)
// 	payoutUC := payoutUC.NewPayoutUC(mockPayoutRepo, log)

// 	mockPayoutRepo.EXPECT().GetOriginalPayoutCache("2XNGAK").Return("").Times(1)
// 	mockPayoutRepo.EXPECT().GetOriginalPayout("2XNGAK").Return("", errors.New("error db")).Times(1)

// 	_, err := payoutUC.GetOriginalPayout("2XNGAK")
// 	if err == nil {
// 		t.Fail()
// 	}
// }

// func TestGetOriginalUrlSuccess(t *testing.T) {
// 	mockCtrl := gomock.NewController(t)
// 	defer mockCtrl.Finish()

// 	mockPayoutRepo := mocks.NewMockPayoutRepository(mockCtrl)
// 	payoutUC := payoutUC.NewPayoutUC(mockPayoutRepo, log)

// 	mockPayoutRepo.EXPECT().GetOriginalPayoutCache("2XNGAK").Return("").Times(1)
// 	mockPayoutRepo.EXPECT().GetOriginalPayout("2XNGAK").Return("https://www.google.com", nil).Times(1)
// 	mockPayoutRepo.EXPECT().SetPayoutCache(&model.Payout{
// 		OriginalPayout: "https://www.google.com",
// 		ShortPayout:    "2XNGAK",
// 	})

// 	oriPayout, err := payoutUC.GetOriginalPayout("2XNGAK")
// 	if err != nil {
// 		t.Fail()
// 	}

// 	if oriPayout != "https://www.google.com" {
// 		t.Fail()
// 	}
// }
