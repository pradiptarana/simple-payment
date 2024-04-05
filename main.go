package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginlogrus "github.com/toorop/gin-logrus"

	dbd "github.com/pradiptarana/simple-payment/internal/db"
	env "github.com/pradiptarana/simple-payment/internal/env"
	limiter "github.com/pradiptarana/simple-payment/internal/limiter"
	logger "github.com/pradiptarana/simple-payment/internal/logger"
	payoutRepo "github.com/pradiptarana/simple-payment/repository/payout"
	payoutTr "github.com/pradiptarana/simple-payment/transport/api/payout"
	payoutUC "github.com/pradiptarana/simple-payment/usecase/payout"
)

func main() {
	log := logger.Init()

	client := http.Client{}
	err := env.LoadEnv()
	if err != nil {
		log.Fatalf("error when load env file")
	}
	db := dbd.NewDBConnection()
	urRepo := payoutRepo.NewPayoutRepository(db, &client)
	payoutUC := payoutUC.NewPayoutUC(urRepo, log)
	payoutTr := payoutTr.NewURLTransport(payoutUC)

	router := gin.Default()
	router.Use(ginlogrus.Logger(log), gin.Recovery())
	r := router.Group("/api/v1")
	r.GET("/account_validation", limiter.Middleware(), payoutTr.AccountValidation)
	r.POST("/payout", limiter.Middleware(), payoutTr.CreatePayout)
	r.POST("/payout/notification", limiter.Middleware(), payoutTr.PayoutNotification)

	router.Run("localhost:8080")
}
