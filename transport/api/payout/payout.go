package payout

import (
	"fmt"
	"net/http"

	"github.com/pradiptarana/simple-payment/model"
	"github.com/pradiptarana/simple-payment/usecase"

	"github.com/gin-gonic/gin"
)

type URLTransport struct {
	usecase.PayoutUsecase
}

func NewURLTransport(uc usecase.PayoutUsecase) *URLTransport {
	return &URLTransport{uc}
}

func (ut *URLTransport) CreatePayout(c *gin.Context) {
	var req *model.PayoutRequest
	if err := c.BindJSON(&req); err != nil {
		fmt.Println("err0", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	payouts, err := ut.PayoutUsecase.Payout(req)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "create payout success", "data": payouts})
	return
}

func (ut *URLTransport) PayoutNotification(c *gin.Context) {
	var req *model.PayoutNotificationRequest
	if err := c.BindJSON(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := ut.PayoutUsecase.PayoutNotification(req)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "notification receive successfully"})
	return
}

func (ut *URLTransport) AccountValidation(c *gin.Context) {
	var req *model.AccountValidationRequest
	if err := c.Bind(&req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	data, err := ut.PayoutUsecase.GetAccountValidation(req)
	if err != nil {
		c.IndentedJSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "ok", "data": data})
	return
}
