package main

import (
	"github.com/acse-sm321/Mogo/paymentApi/handler"
	log "github.com/micro/go-micro/v2/logger"

	paymentApi "github.com/acse-sm321/Mogo/paymentApi/proto/paymentApi"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.paymentApi"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	paymentApi.RegisterPaymentApiHandler(service.Server(), new(handler.PaymentApi))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
